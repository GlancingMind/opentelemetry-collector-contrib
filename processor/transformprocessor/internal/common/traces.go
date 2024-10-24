// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package common // import "github.com/GlancingMind/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/GlancingMind/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/GlancingMind/opentelemetry-collector-contrib/internal/filter/filterottl"
	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlscope"
	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/GlancingMind/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
)

var _ consumer.Traces = &traceStatements{}

type traceStatements struct {
	ottl.StatementSequence[ottlspan.TransformContext]
	expr.BoolExpr[ottlspan.TransformContext]
}

func (t traceStatements) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{
		MutatesData: true,
	}
}

func (t traceStatements) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	for i := 0; i < td.ResourceSpans().Len(); i++ {
		rspans := td.ResourceSpans().At(i)
		for j := 0; j < rspans.ScopeSpans().Len(); j++ {
			sspans := rspans.ScopeSpans().At(j)
			spans := sspans.Spans()
			for k := 0; k < spans.Len(); k++ {
				tCtx := ottlspan.NewTransformContext(spans.At(k), sspans.Scope(), rspans.Resource(), sspans, rspans)
				condition, err := t.BoolExpr.Eval(ctx, tCtx)
				if err != nil {
					return err
				}
				if condition {
					err := t.Execute(ctx, tCtx)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

var _ consumer.Traces = &spanEventStatements{}

type spanEventStatements struct {
	ottl.StatementSequence[ottlspanevent.TransformContext]
	expr.BoolExpr[ottlspanevent.TransformContext]
}

func (s spanEventStatements) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{
		MutatesData: true,
	}
}

func (s spanEventStatements) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	for i := 0; i < td.ResourceSpans().Len(); i++ {
		rspans := td.ResourceSpans().At(i)
		for j := 0; j < rspans.ScopeSpans().Len(); j++ {
			sspans := rspans.ScopeSpans().At(j)
			spans := sspans.Spans()
			for k := 0; k < spans.Len(); k++ {
				span := spans.At(k)
				spanEvents := span.Events()
				for n := 0; n < spanEvents.Len(); n++ {
					tCtx := ottlspanevent.NewTransformContext(spanEvents.At(n), span, sspans.Scope(), rspans.Resource(), sspans, rspans)
					condition, err := s.BoolExpr.Eval(ctx, tCtx)
					if err != nil {
						return err
					}
					if condition {
						err := s.Execute(ctx, tCtx)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

type TraceParserCollection struct {
	parserCollection
	spanParser      ottl.Parser[ottlspan.TransformContext]
	spanEventParser ottl.Parser[ottlspanevent.TransformContext]
}

type TraceParserCollectionOption func(*TraceParserCollection) error

func WithSpanParser(functions map[string]ottl.Factory[ottlspan.TransformContext]) TraceParserCollectionOption {
	return func(tp *TraceParserCollection) error {
		spanParser, err := ottlspan.NewParser(functions, tp.settings)
		if err != nil {
			return err
		}
		tp.spanParser = spanParser
		return nil
	}
}

func WithSpanEventParser(functions map[string]ottl.Factory[ottlspanevent.TransformContext]) TraceParserCollectionOption {
	return func(tp *TraceParserCollection) error {
		spanEventParser, err := ottlspanevent.NewParser(functions, tp.settings)
		if err != nil {
			return err
		}
		tp.spanEventParser = spanEventParser
		return nil
	}
}

func WithTraceErrorMode(errorMode ottl.ErrorMode) TraceParserCollectionOption {
	return func(tp *TraceParserCollection) error {
		tp.errorMode = errorMode
		return nil
	}
}

func NewTraceParserCollection(settings component.TelemetrySettings, options ...TraceParserCollectionOption) (*TraceParserCollection, error) {
	rp, err := ottlresource.NewParser(ResourceFunctions(), settings)
	if err != nil {
		return nil, err
	}
	sp, err := ottlscope.NewParser(ScopeFunctions(), settings)
	if err != nil {
		return nil, err
	}
	tpc := &TraceParserCollection{
		parserCollection: parserCollection{
			settings:       settings,
			resourceParser: rp,
			scopeParser:    sp,
		},
	}

	for _, op := range options {
		err := op(tpc)
		if err != nil {
			return nil, err
		}
	}

	return tpc, nil
}

func (pc TraceParserCollection) ParseContextStatements(contextStatements ContextStatements) (consumer.Traces, error) {
	switch contextStatements.Context {
	case Span:
		parsedStatements, err := pc.spanParser.ParseStatements(contextStatements.Statements)
		if err != nil {
			return nil, err
		}
		globalExpr, errGlobalBoolExpr := parseGlobalExpr(filterottl.NewBoolExprForSpan, contextStatements.Conditions, pc.parserCollection, filterottl.StandardSpanFuncs())
		if errGlobalBoolExpr != nil {
			return nil, errGlobalBoolExpr
		}
		sStatements := ottlspan.NewStatementSequence(parsedStatements, pc.settings, ottlspan.WithStatementSequenceErrorMode(pc.errorMode))
		return traceStatements{sStatements, globalExpr}, nil
	case SpanEvent:
		parsedStatements, err := pc.spanEventParser.ParseStatements(contextStatements.Statements)
		if err != nil {
			return nil, err
		}
		globalExpr, errGlobalBoolExpr := parseGlobalExpr(filterottl.NewBoolExprForSpanEvent, contextStatements.Conditions, pc.parserCollection, filterottl.StandardSpanEventFuncs())
		if errGlobalBoolExpr != nil {
			return nil, errGlobalBoolExpr
		}
		seStatements := ottlspanevent.NewStatementSequence(parsedStatements, pc.settings, ottlspanevent.WithStatementSequenceErrorMode(pc.errorMode))
		return spanEventStatements{seStatements, globalExpr}, nil
	default:
		return pc.parseCommonContextStatements(contextStatements)
	}
}
