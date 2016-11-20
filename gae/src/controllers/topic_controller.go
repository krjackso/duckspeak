package controllers

import (
	"google.golang.org/appengine"
	"net/http"
	"time"

	"models"
	"util/auth"
)

type TopicController struct {
	Router        *DuckRouter
	Authenticator *auth.Authenticator
}

type NewTopicRequest struct {
	Description string    `json:"description"`
	Choices     []string  `json:"choices"`
	ActiveSince time.Time `json:"activeSince"`
	ActiveUntil time.Time `json:"activeUntil"`
}

func (self *TopicController) NewTopic(w http.ResponseWriter, r *http.Request) {
	withAdminAuth(self.Authenticator, w, r, func(adminEmail string) {
		var requestData NewTopicRequest
		if err := parseJson(r, &requestData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if requestData.Description == "" || len(requestData.Choices) != 2 || requestData.ActiveSince.IsZero() || requestData.ActiveUntil.IsZero() {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		ctx := appengine.NewContext(r)
		topic, err := models.NewTopic(ctx, requestData.Description, requestData.Choices, requestData.ActiveSince, requestData.ActiveUntil)
		if err != nil {
			panic(err)
		}

		jsonResponse(w, topic, http.StatusCreated)
	})
}

type GetTopicResponse struct {
	*models.Topic
	Href string `json:"href,omitempty"`
}

func (self *TopicController) GetTopic(w http.ResponseWriter, r *http.Request) {
	withAdminAuth(self.Authenticator, w, r, func(adminEmail string) {
		topicId := self.Router.GetVar(r, "topicId")

		ctx := appengine.NewContext(r)
		topic, err := models.GetTopic(ctx, topicId)
		if err != nil {
			panic(err)
		}

		if topic == nil {
			http.NotFound(w, r)
			return
		}

		response := &GetTopicResponse{
			Topic: topic,
			Href:  self.Router.GetHref(r, "getTopic", "topicId", topic.Id),
		}

		jsonResponse(w, response, http.StatusOK)
	})
}

func (self *TopicController) ListTopics(w http.ResponseWriter, r *http.Request) {
	withAdminAuth(self.Authenticator, w, r, func(adminEmail string) {
		ctx := appengine.NewContext(r)
		topics, err := models.ListTopics(ctx)
		if err != nil {
			panic(err)
		}

		response := make([]*GetTopicResponse, len(topics))
		for index, topic := range topics {
			response[index] = &GetTopicResponse{
				Topic: topic,
				Href:  self.Router.GetHref(r, "getTopic", "topicId", topic.Id),
			}
		}

		jsonResponse(w, response, http.StatusOK)
	})
}

func (self *TopicController) GetDailyTopic(w http.ResponseWriter, r *http.Request) {
	withAccessAuth(self.Authenticator, w, r, func(deviceId string) {
		ctx := appengine.NewContext(r)
		dailyTopic, err := models.GetDailyTopic(ctx)

		if err != nil {
			panic(err)
		}

		if dailyTopic == nil {
			http.NotFound(w, r)
			return
		}

		response := &GetTopicResponse{
			Topic: dailyTopic,
		}

		jsonResponse(w, response, http.StatusCreated)
	})
}
