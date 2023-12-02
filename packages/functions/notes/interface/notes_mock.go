// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package note

import (
	domain "github.com/o-ga09/golang-svelte-app/domain/note"
	"sync"
)

// Ensure, that INotesMock does implement INotes.
// If this is not the case, regenerate this file with moq.
var _ INotes = &INotesMock{}

// INotesMock is a mock implementation of INotes.
//
//	func TestSomethingThatUsesINotes(t *testing.T) {
//
//		// make and configure a mocked INotes
//		mockedINotes := &INotesMock{
//			CreateFunc: func(note domain.Note) error {
//				panic("mock out the Create method")
//			},
//			DeleteFunc: func(userid domain.UserId, noteid domain.NoteId) error {
//				panic("mock out the Delete method")
//			},
//			GetAllFunc: func(id domain.UserId) (domain.Notes, error) {
//				panic("mock out the GetAll method")
//			},
//			GetNoteByIdFunc: func(userId domain.UserId, noteId domain.NoteId) (domain.Note, error) {
//				panic("mock out the GetNoteById method")
//			},
//			UpdateFunc: func(note domain.Note) error {
//				panic("mock out the Update method")
//			},
//		}
//
//		// use mockedINotes in code that requires INotes
//		// and then make assertions.
//
//	}
type INotesMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(note domain.Note) error

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(userid domain.UserId, noteid domain.NoteId) error

	// GetAllFunc mocks the GetAll method.
	GetAllFunc func(id domain.UserId) (domain.Notes, error)

	// GetNoteByIdFunc mocks the GetNoteById method.
	GetNoteByIdFunc func(userId domain.UserId, noteId domain.NoteId) (domain.Note, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(note domain.Note) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Note is the note argument value.
			Note domain.Note
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Userid is the userid argument value.
			Userid domain.UserId
			// Noteid is the noteid argument value.
			Noteid domain.NoteId
		}
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
			// ID is the id argument value.
			ID domain.UserId
		}
		// GetNoteById holds details about calls to the GetNoteById method.
		GetNoteById []struct {
			// UserId is the userId argument value.
			UserId domain.UserId
			// NoteId is the noteId argument value.
			NoteId domain.NoteId
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Note is the note argument value.
			Note domain.Note
		}
	}
	lockCreate      sync.RWMutex
	lockDelete      sync.RWMutex
	lockGetAll      sync.RWMutex
	lockGetNoteById sync.RWMutex
	lockUpdate      sync.RWMutex
}

// Create calls CreateFunc.
func (mock *INotesMock) Create(note domain.Note) error {
	if mock.CreateFunc == nil {
		panic("INotesMock.CreateFunc: method is nil but INotes.Create was just called")
	}
	callInfo := struct {
		Note domain.Note
	}{
		Note: note,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(note)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedINotes.CreateCalls())
func (mock *INotesMock) CreateCalls() []struct {
	Note domain.Note
} {
	var calls []struct {
		Note domain.Note
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *INotesMock) Delete(userid domain.UserId, noteid domain.NoteId) error {
	if mock.DeleteFunc == nil {
		panic("INotesMock.DeleteFunc: method is nil but INotes.Delete was just called")
	}
	callInfo := struct {
		Userid domain.UserId
		Noteid domain.NoteId
	}{
		Userid: userid,
		Noteid: noteid,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(userid, noteid)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedINotes.DeleteCalls())
func (mock *INotesMock) DeleteCalls() []struct {
	Userid domain.UserId
	Noteid domain.NoteId
} {
	var calls []struct {
		Userid domain.UserId
		Noteid domain.NoteId
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// GetAll calls GetAllFunc.
func (mock *INotesMock) GetAll(id domain.UserId) (domain.Notes, error) {
	if mock.GetAllFunc == nil {
		panic("INotesMock.GetAllFunc: method is nil but INotes.GetAll was just called")
	}
	callInfo := struct {
		ID domain.UserId
	}{
		ID: id,
	}
	mock.lockGetAll.Lock()
	mock.calls.GetAll = append(mock.calls.GetAll, callInfo)
	mock.lockGetAll.Unlock()
	return mock.GetAllFunc(id)
}

// GetAllCalls gets all the calls that were made to GetAll.
// Check the length with:
//
//	len(mockedINotes.GetAllCalls())
func (mock *INotesMock) GetAllCalls() []struct {
	ID domain.UserId
} {
	var calls []struct {
		ID domain.UserId
	}
	mock.lockGetAll.RLock()
	calls = mock.calls.GetAll
	mock.lockGetAll.RUnlock()
	return calls
}

// GetNoteById calls GetNoteByIdFunc.
func (mock *INotesMock) GetNoteById(userId domain.UserId, noteId domain.NoteId) (domain.Note, error) {
	if mock.GetNoteByIdFunc == nil {
		panic("INotesMock.GetNoteByIdFunc: method is nil but INotes.GetNoteById was just called")
	}
	callInfo := struct {
		UserId domain.UserId
		NoteId domain.NoteId
	}{
		UserId: userId,
		NoteId: noteId,
	}
	mock.lockGetNoteById.Lock()
	mock.calls.GetNoteById = append(mock.calls.GetNoteById, callInfo)
	mock.lockGetNoteById.Unlock()
	return mock.GetNoteByIdFunc(userId, noteId)
}

// GetNoteByIdCalls gets all the calls that were made to GetNoteById.
// Check the length with:
//
//	len(mockedINotes.GetNoteByIdCalls())
func (mock *INotesMock) GetNoteByIdCalls() []struct {
	UserId domain.UserId
	NoteId domain.NoteId
} {
	var calls []struct {
		UserId domain.UserId
		NoteId domain.NoteId
	}
	mock.lockGetNoteById.RLock()
	calls = mock.calls.GetNoteById
	mock.lockGetNoteById.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *INotesMock) Update(note domain.Note) error {
	if mock.UpdateFunc == nil {
		panic("INotesMock.UpdateFunc: method is nil but INotes.Update was just called")
	}
	callInfo := struct {
		Note domain.Note
	}{
		Note: note,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(note)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//
//	len(mockedINotes.UpdateCalls())
func (mock *INotesMock) UpdateCalls() []struct {
	Note domain.Note
} {
	var calls []struct {
		Note domain.Note
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
