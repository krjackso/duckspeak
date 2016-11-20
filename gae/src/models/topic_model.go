package models

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"time"
)

type Topic struct {
	Id          string    `json:"id" datastore:"-"`
	Created     time.Time `json:"created"`
	ActiveSince time.Time `json:"activeSince"`
	ActiveUntil time.Time `json:"activeUntil"`
	Description string    `json:"description" datastore:",noindex"`
	Choices     []string  `json:"choices" datastore:",noindex"`
}

func topicKey(ctx context.Context, topicId string) *datastore.Key {
	return datastore.NewKey(ctx, "Topic", topicId, 0, nil)
}

func GetTopic(ctx context.Context, topicId string) (*Topic, error) {
	key := topicKey(ctx, topicId)

	topic := &Topic{}
	err := datastore.Get(ctx, key, topic)

	switch err {
	case datastore.ErrNoSuchEntity:
		return nil, nil
	case nil:
		break
	default:
		return nil, err
	}

	topic.Id = topicId
	return topic, nil
}

func GetDailyTopic(ctx context.Context) (*Topic, error) {
	now := time.Now()
	query := datastore.NewQuery("Topic").Filter("ActiveUntil >", now).Order("-ActiveUntil").Limit(1)

	iter := query.Run(ctx)
	var topic Topic

	key, err := iter.Next(&topic)

	if err == datastore.Done {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	// Check ActiveSince because datastore only allows one inequality filter
	if topic.ActiveSince.After(now) {
		return nil, nil
	}

	topic.Id = key.StringID()
	return &topic, nil
}

func ListTopics(ctx context.Context) ([]*Topic, error) {
	query := datastore.NewQuery("Topic").Order("-ActiveUntil")
	var topics []*Topic
	keys, err := query.GetAll(ctx, &topics)

	if err != nil {
		return nil, err
	}

	for i, topic := range topics {
		topic.Id = keys[i].StringID()
	}

	return topics, nil
}

func NewTopic(ctx context.Context, description string, choices []string, since, until time.Time) (*Topic, error) {
	newId := uuid.NewV4().String()
	key := topicKey(ctx, newId)

	topic := &Topic{
		Id:          newId,
		Created:     time.Now().UTC(),
		ActiveSince: since.UTC(),
		ActiveUntil: until.UTC(),
		Description: description,
		Choices:     choices,
	}

	key, err := datastore.Put(ctx, key, topic)
	if err != nil {
		return nil, err
	}

	return topic, nil
}
