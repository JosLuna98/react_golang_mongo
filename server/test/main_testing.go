package main_testing

import (
	"server/models"
	"server/middleware"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test() {
	testInsertOneTask()
	testGetAllTask()
	testCompleteTask()
	testUndoTask()
	testDeleteOneTask()
	fmt.Println("testing done!")
}

var testing_task string

func testInsertOneTask() {
	var initial_task = models.ToDoList{
		Description: "testing_task",
	}

	var result = middleware.TinsertOneTask(initial_task)

	if result["error"] == true {
		fmt.Println("can't insert one task")
	}

	fmt.Println("testing: task created")
}

func testGetAllTask() {
	var result = middleware.TgetAllTask()

	if result["error"] == true {
		fmt.Println("can't get all task")
	}

	testing_task = result["result"].([]primitive.M)[0]["_id"].(primitive.ObjectID).Hex()
	fmt.Println("testing: task gotten")
}

func testCompleteTask() {
	var result = middleware.TcompleteTask(testing_task)

	if result["error"] == true {
		fmt.Println("can't complete a task")
	}

	fmt.Println("testing: task completed")
}

func testUndoTask() {
	var result = middleware.TundoTask(testing_task)

	if result["error"] == true {
		fmt.Println("can't undo a task")
	}

	fmt.Println("testing: task undone")
}

func testDeleteOneTask() {
	var result = middleware.TdeleteOneTask(testing_task)

	if result["error"] == true {
		fmt.Println("can't delete one task")
	}

	fmt.Println("testing: task deleted")
}
