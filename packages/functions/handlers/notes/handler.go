package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	domain "github.com/o-ga09/golang-svelte-app/domain/note"
	note "github.com/o-ga09/golang-svelte-app/notes/interface"
)

type NoteHandler struct {
	Notes note.INotes
}

type Note struct {
	UserId     string `json:"user_id,omitempty"`
	NoteId     string `json:"note_id,omitempty"`
	Contents   string `json:"contents,omitempty"`
	Image_url  string `json:"image_url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
}

func New(note note.INotes) *NoteHandler {
	return &NoteHandler{Notes: note}
}

func(h *NoteHandler) GetAll(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	dummyUserId, err := uuid.NewRandom()
	reqUserId := dummyUserId.String()
	userId := domain.UserId{Value: reqUserId}
	notes, err := h.Notes.GetAll(userId)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
		}, err
	}


	body := []Note{}
	for _, note := range notes {
		b := Note{
			UserId: note.UserId.Value,
			NoteId: note.NoteId.Value,
			Contents: note.Contents.Value,
			Image_url: note.Image_url.Value,
			Created_at: note.Created_at.Value.String(),
			Updated_at: note.Updated_at.Value.String(),
		}

		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
			}, err
		}
		body = append(body,b)
	}
	
	res, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body: string(res),
	}, nil
}
func(h *NoteHandler) GetNoteById(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	reqNoteId := req.PathParameters["id"]
	noteId := domain.NoteId{Value: reqNoteId}
	dummyUserId, err := uuid.NewRandom()
	reqUserId := dummyUserId.String()
	userId := domain.UserId{Value: reqUserId}

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	note, err := h.Notes.GetNoteById(userId,noteId)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
		}, err
	}
	
	res := Note{
		UserId: note.UserId.Value,
		NoteId: note.NoteId.Value,
		Contents: note.Contents.Value,
		Image_url: note.Image_url.Value,
		Created_at: note.Created_at.Value.String(),
		Updated_at: note.Updated_at.Value.String(),
	}

	body, err := json.Marshal(res)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body: string(body),
	}, nil
}

func(h *NoteHandler) Create(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	reqNoteId := req.PathParameters["id"]
	noteId := domain.NoteId{Value: reqNoteId}
	dummyUserId, err := uuid.NewRandom()
	reqUserId := dummyUserId.String()
	userId := domain.UserId{Value: reqUserId}
	reqBody := Note{}
	if err := json.Unmarshal([]byte(req.Body),&reqBody); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusCreated,
		}, err
	}
	body := domain.Note{
		UserId: userId,
		NoteId: noteId,
		Contents: domain.Contents{Value: reqBody.Contents},
		Image_url: domain.Image_url{Value: reqBody.Image_url},
	}
	err = h.Notes.Create(body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
	}, nil
}
func(h *NoteHandler) Update(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	reqNoteId := req.PathParameters["id"]
	noteId := domain.NoteId{Value: reqNoteId}
	dummyUserId, err := uuid.NewRandom()
	reqUserId := dummyUserId.String()
	userId := domain.UserId{Value: reqUserId}

	reqBody := Note{}
	if err := json.Unmarshal([]byte(req.Body),&reqBody); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusCreated,
		}, err
	}
	body := domain.Note{
		UserId: userId,
		NoteId: noteId,
		Contents: domain.Contents{Value: reqBody.Contents},
		Image_url: domain.Image_url{Value: reqBody.Image_url},
	}

	err = h.Notes.Update(body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
	}, nil
}
func(h *NoteHandler) Delete(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	reqNoteId := req.PathParameters["id"]
	noteId := domain.NoteId{Value: reqNoteId}
	dummyUserId, err := uuid.NewRandom()
	reqUserId := dummyUserId.String()
	userId := domain.UserId{Value: reqUserId}

	err = h.Notes.Delete(userId,noteId)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}