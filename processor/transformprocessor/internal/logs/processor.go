// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/GlancingMind/opentelemetry-collector-contrib/processor/transformprocessor/internal/logs"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/GlancingMind/opentelemetry-collector-contrib/internal/pdatautil"
	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/GlancingMind/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"
)

type Processor struct {
	contexts []consumer.Logs
	logger   *zap.Logger
	flatMode bool
}

func NewProcessor(contextStatements []common.ContextStatements, errorMode ottl.ErrorMode, flatMode bool, settings component.TelemetrySettings) (*Processor, error) {
	pc, err := common.NewLogParserCollection(settings, common.WithLogParser(LogFunctions()), common.WithLogErrorMode(errorMode))
	if err != nil {
		return nil, err
	}

	contexts := make([]consumer.Logs, len(contextStatements))
	var errors error
	for i, cs := range contextStatements {
		context, err := pc.ParseContextStatements(cs)
		if err != nil {
			errors = multierr.Append(errors, err)
		}
		contexts[i] = context
	}

	if errors != nil {
		return nil, errors
	}

	return &Processor{
		contexts: contexts,
		logger:   settings.Logger,
		flatMode: flatMode,
	}, nil
}

func (p *Processor) ProcessLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	if p.flatMode {
		pdatautil.FlattenLogs(ld.ResourceLogs())
		defer pdatautil.GroupByResourceLogs(ld.ResourceLogs())
	}
	for _, c := range p.contexts {
		err := c.ConsumeLogs(ctx, ld)
		if err != nil {
			p.logger.Error("failed processing logs", zap.Error(err))
			return ld, err
		}
	}
	return ld, nil
}
