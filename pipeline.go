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

func (p *Pipeline[T, R]) AddFields(fields bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$addFields", fields}})
	return p
}

func (p *Pipeline[T, R]) Bucket(bucket bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$bucket", bucket}})
	return p
}

func (p *Pipeline[T, R]) BucketAuto(bucket bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$bucketAuto", bucket}})
	return p
}

func (p *Pipeline[T, R]) ChangeStream(changeStream bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$changeStream", changeStream}})
	return p
}

func (p *Pipeline[T, R]) ChangeStreamSplitLargeEvents(changeStream bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$changeStream", changeStream}})
	return p
}

func (p *Pipeline[T, R]) CollStats(collStats bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$collStats", collStats}})
	return p
}

func (p *Pipeline[T, R]) Count(count bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$count", count}})
	return p
}

func (p *Pipeline[T, R]) CurrentOp(currentOp bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$currentOp", currentOp}})
	return p
}

func (p *Pipeline[T, R]) Densify(densify bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$densify", densify}})
	return p
}

func (p *Pipeline[T, R]) Documents(documents bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$documents", documents}})
	return p
}

func (p *Pipeline[T, R]) Facet(facet bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$facet", facet}})
	return p
}

func (p *Pipeline[T, R]) Fill(fill bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$fill", fill}})
	return p
}

func (p *Pipeline[T, R]) GeoNear(geoNear bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$geoNear", geoNear}})
	return p
}

func (p *Pipeline[T, R]) GraphLookup(graphLookup bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$graphLookup", graphLookup}})
	return p
}

func (p *Pipeline[T, R]) Group(group bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$group", group}})
	return p
}

func (p *Pipeline[T, R]) IndexStats(indexStats bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$indexStats", indexStats}})
	return p
}

func (p *Pipeline[T, R]) Limit(limit bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$limit", limit}})
	return p
}

func (p *Pipeline[T, R]) ListLocalSessions(listLocalSessions bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$listLocalSessions", listLocalSessions}})
	return p
}

func (p *Pipeline[T, R]) ListSampledQueries(listSampledQueries bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$listSampledQueries", listSampledQueries}})
	return p
}

func (p *Pipeline[T, R]) ListSearchIndexes(listSearchIndexes bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$listSearchIndexes", listSearchIndexes}})
	return p
}

func (p *Pipeline[T, R]) ListSessions(listSessions bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$listSessions", listSessions}})
	return p
}

func (p *Pipeline[T, R]) Lookup(lookup bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$lookup", lookup}})
	return p
}

func (p *Pipeline[T, R]) Match(filter bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$match", filter}})
	return p
}

func (p *Pipeline[T, R]) Merge(merge bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$merge", merge}})
	return p
}

func (p *Pipeline[T, R]) Out(out bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$out", out}})
	return p
}

func (p *Pipeline[T, R]) PlanCacheStats(planCacheStats bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$planCacheStats", planCacheStats}})
	return p
}

func (p *Pipeline[T, R]) Project(project bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$project", project}})
	return p
}

func (p *Pipeline[T, R]) Redact(redact bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$redact", redact}})
	return p
}

func (p *Pipeline[T, R]) ReplaceRoot(replaceRoot bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$replaceRoot", replaceRoot}})
	return p
}

func (p *Pipeline[T, R]) ReplaceWith(replaceWith bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$replaceWith", replaceWith}})
	return p
}

func (p *Pipeline[T, R]) Sample(sample bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$sample", sample}})
	return p
}

func (p *Pipeline[T, R]) Search(search bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$search", search}})
	return p
}

func (p *Pipeline[T, R]) SearchMeta(searchMeta bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$searchMeta", searchMeta}})
	return p
}

func (p *Pipeline[T, R]) Set(set bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$set", set}})
	return p
}

func (p *Pipeline[T, R]) SetWindowFields(setWindowFields bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$setWindowFields", setWindowFields}})
	return p
}

func (p *Pipeline[T, R]) SharedDataDistribution(sharedDataDistribution bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$sharedDataDistribution", sharedDataDistribution}})
	return p
}

func (p *Pipeline[T, R]) Skip(skip bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$skip", skip}})
	return p
}

func (p *Pipeline[T, R]) Sort(sort bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$sort", sort}})
	return p
}

func (p *Pipeline[T, R]) SortByCount(sortByCount bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$sortByCount", sortByCount}})
	return p
}

func (p *Pipeline[T, R]) UnionWith(unionWith bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$unionWith", unionWith}})
	return p
}

func (p *Pipeline[T, R]) Unset(unset bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$unset", unset}})
	return p
}

func (p *Pipeline[T, R]) Unwind(unwind bson.M) *Pipeline[T, R] {
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
