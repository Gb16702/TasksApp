package routes

import (
	"fmt"
	"strconv"
	"todoapp/core/utils"
	"todoapp/core/utils/validation"
	"todoapp/database"
	"todoapp/database/models"

	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type Task struct {
	Uid 		uint 	`json:"uid"`
	Name 		string 	`json:"name"`
	Done 		bool
}

type TasksResponse struct {
    Tasks []Task `json:"tasks"`
}

func HandleTasksRoutes(app *fiber.App) {
	app.Get("/api/tasks/:uid", func (c *fiber.Ctx) error {
		if c.Method() != fiber.MethodGet {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorMethod"), 405)
		};

		uid := c.Params("uid");

		if uid == "" {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalidParams"), 400)
		}

		id, err := strconv.Atoi(uid);

		if err != nil {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalidParams"), 400)
		}

		tasks, err := getTasksByUserId(id);

		if err != nil {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalidParams"), 400)
		}

		response := TasksResponse{Tasks: tasks}
        responseToJson, err := json.Marshal(response)

		if err != nil {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalidParams"), 400)
		}

		return utils.GenerateResponse(c, string(responseToJson), 200)
	})

	app.Post("/api/tasks", func(c *fiber.Ctx) error {
		if c.Method() != fiber.MethodPost {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorMethod"), 405)
		}

		var newTask Task

		fmt.Println(c.Body())

		if err := c.BodyParser(&newTask); err != nil {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalid"), 400)
		}

		requiredFields := []struct{
			name 				string
			err 				string
		} {
			{name: "uid", err: "L'identifiant de l'utilisateur est requis"},
			{name: "name", err: "Le titre est requis"},
		}

		for _, field := range requiredFields {
			switch field.name {
			case "uid":
				isEmpty, err := validation.IsFieldEmpty(c, strconv.FormatUint(uint64(newTask.Uid), 10), field.err)
				if isEmpty {
					return err
				}
		 	case "name":
		 		isEmpty, err := validation.IsFieldEmpty(c, newTask.Name, field.err)
		 		if isEmpty {
					return err
				}
			}
		}

		task := models.Task{
			Uid: newTask.Uid,
			Name: newTask.Name,
			Done: false,
		}

		if result := database.DB.Create(&task); result.Error != nil {
			return utils.GenerateResponse(c, "Une erreur est survenue lors de la création de la tâche", 500)
		}

		return utils.GenerateResponse(c, "Tâche créée avec succès", 200)
	})
}

func getTasksByUserId(userID int) ([]Task, error) {
	var tasks []Task;
	result := database.DB.Where("uid = ?", userID).Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}
