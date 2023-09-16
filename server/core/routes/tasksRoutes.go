package routes

import (
	"strconv"
	"todoapp/core/utils"
	"todoapp/core/utils/validation"
	"todoapp/database"
	"todoapp/database/models"

	"github.com/gofiber/fiber/v2"
)

type Task struct {
	Id 			uint 	`json:"id"`
	Uid 		uint 	`json:"uid"`
	Name 		string 	`json:"name"`
	Done 		bool	`json:"done"`

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

		if err != nil {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalidParams"), 400)
		}

		return c.Status(200).JSON(fiber.Map{
			"tasks": tasks,
		})
	})

	app.Post("/api/tasks", func(c *fiber.Ctx) error {
		if c.Method() != fiber.MethodPost {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorMethod"), 405)
		}

		var newTask Task

		if err := c.BodyParser(&newTask); err != nil {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalid"), 400)
		}

		var userTasks []models.Task
		if result := database.DB.Where("uid = ?", newTask.Uid).Find(&userTasks); result.Error != nil {
			return utils.GenerateResponse(c, "Une erreur est survenue lors de la création de la tâche", 500)
		}

		if len(userTasks) >= 10 {
			return utils.GenerateResponse(c, "Tu as atteint le maximum de tâches", 400)
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

	app.Post("/api/tasks/name", func(c *fiber.Ctx) error {
		if c.Method() != fiber.MethodPost {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorMethod"), 405)
		}

		var newTask Task
		if err := c.BodyParser(&newTask); err != nil {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalid"), 400)
		}

		if result := database.DB.Model(&models.Task{}).Where("id = ?", newTask.Id).Update("name", newTask.Name); result.Error != nil {
			return utils.GenerateResponse(c, "Une erreur est survenue lors de la mise à jour de la tâche", 500)
		}

		return utils.GenerateResponse(c, "Tâche mise à jour avec succès", 200)
	})

	app.Delete("/api/tasks/:id", func(c *fiber.Ctx) error {
		if c.Method() != fiber.MethodDelete {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorMethod"), 405)
		}

		id := c.Params("id")
		if id == "" {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalidParams"), 400)
		}

		taskId, err := strconv.Atoi(id)

		if err != nil {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalidParams"), 400)
		}

		if result := database.DB.Unscoped().Where("id = ?", taskId).Delete(&models.Task{}); result.Error != nil {
			return utils.GenerateResponse(c, "Une erreur est survenue lors de la suppression de la tâche", 500)
		};

		return utils.GenerateResponse(c, "Tâche supprimée avec succès", 200);
	})

	app.Patch("/api/tasks/:id", func(c *fiber.Ctx) error  {

		if c.Method() != fiber.MethodPatch {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorMethod"), 405)
		}

		id := c.Params("id")

		if id == "" {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalidParams"), 400)
		}

		taskId, err := strconv.Atoi(id)

		if err != nil {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalidParams"), 400)
		}

		task := models.Task{}
		if result := database.DB.Model(&models.Task{}).Where("id = ?", taskId).Find(&task).Update("done", !task.Done); result.Error != nil {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalidParams"), 400)
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Tâche mise à jour avec succès",
			"task" : task,
		})
	})

	app.Delete("/api/tasks/clear/:id", func(c *fiber.Ctx) error {

		if c.Method() != fiber.MethodDelete {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorMethod"), 405);
		}

		id := c.Params("id")

		if id == "" {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalidParams"), 400)
		}

		userId, err := strconv.Atoi(id);

		if err != nil {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalidParams"), 400)
		}

		if result := database.DB.Unscoped().Where("uid = ?", userId).Delete(&models.Task{}); result.Error != nil {
			return utils.GenerateResponse(c, "Une erreur est survenue lors de la suppression des tâches", 500)
		}

		return utils.GenerateResponse(c, "Tâches supprimées avec succès", 200)
 	})

}

func getTasksByUserId(userID int) ([]Task, error) {
	var tasks []Task;
	result := database.DB.Where("uid = ?", userID).Find(&tasks).Order("created_at DESC")

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}
