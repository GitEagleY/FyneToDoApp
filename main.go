package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Task struct {
	ID          uint
	Title       string
	Description string
}

func main() {
	///////////////////APP minimal init///////////////////////
	app := app.New()                           //create app
	app.Settings().SetTheme(theme.DarkTheme()) //set dark theme

	window := app.NewWindow("TODO app") //create window for app
	window.CenterOnScreen()             //center window
	//////////////////////////////////////////////////////////

	var tasks []Task //tasks array

	//////////////////DB LOGIC/////////////////////////////////////////////////
	DB, _ := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{}) //initialize DB
	DB.AutoMigrate(&Task{})                                    //migrate db
	DB.Find(&tasks)                                            //find all task in db
	//////////////////////////////////////////////////////////////////////////

	var tasksList *widget.List        //initialize tasks list
	var CreateTaskBox *fyne.Container //initialize create content

	/////////////////////////////////INPUT LABELS//////////////////////////////////////////
	titleEntry := widget.NewEntry()                              //initialize title entry
	titleEntry.SetPlaceHolder("Task title...")                   //Title to entry
	DescriptionEntry := widget.NewMultiLineEntry()               //init descr entry
	DescriptionEntry.SetPlaceHolder("Description for your Task") //description to entry
	//////////////////////////////////////////////////////////////////////////////////////

	////////////////////////////////////IDK/////////////////////////////////
	MainBoxWithTasks := container.NewVBox() //initialize main box where to show tasks
	window.SetContent(MainBoxWithTasks)     //set content
	////////////////////////////////////////////////////////////////////////

	ButtonBackAndSave := widget.NewButton("back", func() {

		/////////////////////////////ADDING TASK TO DB if title field not empty///////////////////////////////
		if titleEntry.Text != "" {
			task := Task{Title: titleEntry.Text,
				Description: DescriptionEntry.Text}
			DB.Create(&task)
			DB.Find(&task)

		}
		/////////////////////////////////////////////////////////////////////////////////////////////////
		////////////////////Clear and update inputs////////////////////
		titleEntry.Text = ""
		titleEntry.Refresh()
		DescriptionEntry.Text = ""
		DescriptionEntry.Refresh()
		/////////////////////////////////////////////////////////

		window.SetContent(MainBoxWithTasks)

		DB.Find(&tasks) ////refresh tasks

		tasksList.UnselectAll()

	})

	////////////////////////init BAR with input////////////////////////////
	TopBarCreatingTask := container.NewHBox(layout.NewSpacer(), canvas.NewText("Task:", color.White),
		layout.NewSpacer(), ButtonBackAndSave,
	)
	/////////////////////////////////////////////////////////

	///////////////////////////////init TASKS LIST////////////////////////////////////
	tasksList = widget.NewList(func() int {
		return len(tasks)
	},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(tasks[lii].Title)
		})
	////////////////////////////////////////////////////////////////////////////////

	//////////////////////////////////////SET CREATE TASK BOX////////////////////////////////////
	CreateTaskBox = container.NewVBox(TopBarCreatingTask, container.NewVBox(titleEntry, DescriptionEntry))
	/////////////////////////////////////////////////////////////////////////////////////////////////

	/////////////////////////////OPEN TASK when click/////////////////////////////////
	tasksList.OnSelected = func(id widget.ListItemID) {
		detailBar := container.NewHBox(
			layout.NewSpacer(),
			canvas.NewText(fmt.Sprintf("%s", tasks[id].Title), color.White),
			layout.NewSpacer(),
			widget.NewButton("back", func() {
				window.SetContent(MainBoxWithTasks)
				tasksList.Unselect(id)
				tasksList.UnselectAll()
			}),
		)
		/////////////////////////////DELETE///////////////////////////////////////////////
		DELETEbutton := widget.NewButton("DELETE", func() {

			DB.Delete(&Task{}, "Id = ?", tasks[id].ID)
			DB.Find(&tasks)

			window.SetContent(MainBoxWithTasks)
		})
		//////////////////////////////////////////////////////////////////////////////

		/////////////////////////////BUTTON TO ENTER EDIT MODE////////////////////////////////////////
		EDITbutton := widget.NewButton("EDIT", func() {

			editBar := container.NewHBox(

				canvas.NewText(fmt.Sprintf("Editing \"%s\"", tasks[id].Title), color.White),
				layout.NewSpacer(),
				widget.NewButton("back", func() {
					window.SetContent(MainBoxWithTasks)
					tasksList.Unselect(id)
				}),
			)

			///////////////EDIT TEXT/////////////////
			editTitle := widget.NewEntry()
			editTitle.SetText(tasks[id].Title)
			editDescription := widget.NewMultiLineEntry()
			editDescription.SetText(tasks[id].Description)
			////////////////////////////////////////////////

			SaveAfterEdit := widget.NewButton("Save", func() {
				DB.Find(
					&Task{}, "Id = ? ", tasks[id].ID,
				).Updates(
					Task{
						Title:       editTitle.Text,
						Description: editDescription.Text,
					},
				)
				DB.Find(&tasks)
				window.SetContent(MainBoxWithTasks)
				tasksList.UnselectAll()
			})

			SetEditContent := container.NewVBox(
				editBar,
				canvas.NewLine(color.White),
				editTitle,
				editDescription,
				SaveAfterEdit,
			)
			window.SetContent(SetEditContent)
		})
		//////////////////////////////////////////////////////////////////////////////////////////////////

		/////////////////////BUTTONS BOX//////////////////
		buttonsBox := container.NewHBox(
			layout.NewSpacer(),
			DELETEbutton,
			EDITbutton,
			layout.NewSpacer(),
		)
		/////////////////////////////////////////////////

		taskTitle := widget.NewLabel(tasks[id].Description)
		taskTitle.TextStyle = fyne.TextStyle{Bold: true}
		taskDescription := widget.NewLabel(tasks[id].Description)
		taskDescription.TextStyle = fyne.TextStyle{Bold: true}
		taskDescription.Wrapping = fyne.TextWrapBreak
		detailsVbox := container.NewVBox(
			detailBar,
			canvas.NewLine(color.White),
			taskTitle,

			buttonsBox,
		)
		tasksList.Unselect(id)
		tasksList.UnselectAll()
		window.SetContent(detailsVbox)

	}

	///////////////////////////INIT TASK SCROLL//////////////
	taskScroll := container.NewScroll(tasksList)
	taskScroll.SetMinSize(fyne.NewSize(300, 400))
	//////////////////////////////////////////////////////////

	///////////////////////////////////INIT AND SET TASKBAR////////////////////
	TaskCreateBar := container.NewHBox(
		layout.NewSpacer(),
		canvas.NewText("Tasks", color.White),
		layout.NewSpacer(),
		widget.NewButton("Add new one", func() {
			window.SetContent(CreateTaskBox)
		}),
	)
	////////////////////////////////////////////////////////////////////
	MainBoxWithTasks = container.NewVBox(TaskCreateBar,
		canvas.NewLine(color.White),

		taskScroll)

	window.Resize(fyne.NewSize(500, 500))
	window.SetContent(MainBoxWithTasks)
	window.Show()
	app.Run()

}
