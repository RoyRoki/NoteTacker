package di

import (
	"database/sql"
	"note-tracker/modules/notes/data/datasource"
	"note-tracker/modules/notes/data/repository_impl"
	"note-tracker/modules/notes/domain/usecase"
	"note-tracker/modules/notes/presentation/controller"
	"note-tracker/modules/notes/presentation/router"

	"github.com/gorilla/mux"
)

func RegisterNoteModule(db *sql.DB, r *mux.Router) {
	notesDataSource := datasource.NewNoteDataSource(db)
	notesRepo := repository_impl.NewNoteRepository(notesDataSource)
	notesUsecase := usecase.NewNoteUsecase(notesRepo)
	notesController := controller.NewNoteController(notesUsecase)

	router.RegisterNoteRoutes(r, notesController)
}
