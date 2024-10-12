package colt

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type Pipeline[T Document, R any] struct {
	collection *mongo.Collection
	pipeline   mongo.Pipeline
}

func (p *Pipeline[T, R]) AddFields(fields interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$addFields", fields}})
	return p
}

func (p *Pipeline[T, R]) Bucket(bucket interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$bucket", bucket}})
	return p
}

func (p *Pipeline[T, R]) BucketAuto(bucket interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$bucketAuto", bucket}})
	return p
}

func (p *Pipeline[T, R]) ChangeStream(changeStream interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$changeStream", changeStream}})
	return p
}

func (p *Pipeline[T, R]) ChangeStreamSplitLargeEvents(changeStream interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$changeStream", changeStream}})
	return p
}

func (p *Pipeline[T, R]) CollStats(collStats interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$collStats", collStats}})
	return p
}

func (p *Pipeline[T, R]) Count(count interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$count", count}})
	return p
}

func (p *Pipeline[T, R]) CurrentOp(currentOp interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$currentOp", currentOp}})
	return p
}

func (p *Pipeline[T, R]) Densify(densify interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$densify", densify}})
	return p
}

func (p *Pipeline[T, R]) Documents(documents interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$documents", documents}})
	return p
}

func (p *Pipeline[T, R]) Facet(facet interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$facet", facet}})
	return p
}

func (p *Pipeline[T, R]) Fill(fill interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$fill", fill}})
	return p
}

func (p *Pipeline[T, R]) GeoNear(geoNear interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$geoNear", geoNear}})
	return p
}

func (p *Pipeline[T, R]) GraphLookup(graphLookup interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$graphLookup", graphLookup}})
	return p
}

func (p *Pipeline[T, R]) Group(group interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$group", group}})
	return p
}

func (p *Pipeline[T, R]) IndexStats(indexStats interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$indexStats", indexStats}})
	return p
}

func (p *Pipeline[T, R]) Limit(limit interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$limit", limit}})
	return p
}

func (p *Pipeline[T, R]) ListLocalSessions(listLocalSessions interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$listLocalSessions", listLocalSessions}})
	return p
}

func (p *Pipeline[T, R]) ListSampledQueries(listSampledQueries interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$listSampledQueries", listSampledQueries}})
	return p
}

func (p *Pipeline[T, R]) ListSearchIndexes(listSearchIndexes interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$listSearchIndexes", listSearchIndexes}})
	return p
}

func (p *Pipeline[T, R]) ListSessions(listSessions interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$listSessions", listSessions}})
	return p
}

func (p *Pipeline[T, R]) Lookup(lookup interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$lookup", lookup}})
	return p
}

func (p *Pipeline[T, R]) Match(filter interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$match", filter}})
	return p
}

func (p *Pipeline[T, R]) Merge(merge interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$merge", merge}})
	return p
}

func (p *Pipeline[T, R]) Out(out interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$out", out}})
	return p
}

func (p *Pipeline[T, R]) PlanCacheStats(planCacheStats interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$planCacheStats", planCacheStats}})
	return p
}

func (p *Pipeline[T, R]) Project(project interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$project", project}})
	return p
}

func (p *Pipeline[T, R]) Redact(redact interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$redact", redact}})
	return p
}

func (p *Pipeline[T, R]) ReplaceRoot(replaceRoot interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$replaceRoot", replaceRoot}})
	return p
}

func (p *Pipeline[T, R]) ReplaceWith(replaceWith interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$replaceWith", replaceWith}})
	return p
}

func (p *Pipeline[T, R]) Sample(sample interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$sample", sample}})
	return p
}

func (p *Pipeline[T, R]) Search(search interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$search", search}})
	return p
}

func (p *Pipeline[T, R]) SearchMeta(searchMeta interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$searchMeta", searchMeta}})
	return p
}

func (p *Pipeline[T, R]) Set(set interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$set", set}})
	return p
}

func (p *Pipeline[T, R]) SetWindowFields(setWindowFields interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$setWindowFields", setWindowFields}})
	return p
}

func (p *Pipeline[T, R]) SharedDataDistribution(sharedDataDistribution interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$sharedDataDistribution", sharedDataDistribution}})
	return p
}

func (p *Pipeline[T, R]) Skip(skip interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$skip", skip}})
	return p
}

func (p *Pipeline[T, R]) Sort(sort interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$sort", sort}})
	return p
}

func (p *Pipeline[T, R]) SortByCount(sortByCount interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$sortByCount", sortByCount}})
	return p
}

func (p *Pipeline[T, R]) UnionWith(unionWith interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$unionWith", unionWith}})
	return p
}

func (p *Pipeline[T, R]) Unset(unset interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$unset", unset}})
	return p
}

func (p *Pipeline[T, R]) Unwind(unwind interface{}) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$unwind", unwind}})
	return p
}

func (p *Pipeline[T, R]) AppendStage(stage bson.D) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, stage)
	return p
}

func (p *Pipeline[T, R]) Run() (Cursor[R], error) {
	c, err := p.collection.Aggregate(DefaultContext(), p.pipeline)
	if err != nil {
		return nil, err
	}
	return &cursor[R]{&sync.Mutex{}, nil, nil, false, DefaultContext(), c}, nil
}
