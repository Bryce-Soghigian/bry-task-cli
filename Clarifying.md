# Task Tool I can use to make note of everything i did at microsoft and learn go
We want to create an executable binary that will allow us to
ask the cli to do something
the cli responds back
and shuts down

```
Features
Create Task
task create -t "Learn about Exported names" -d "Research what exported names are in golang"

Complete a task
task complete <task_id> || <task_title>

Abort a task
task abort <task_id> || <task_title>

Task List Commands

task list -a
task list -c 
task list -d 03/31/2000
-------------------------------------------
flags| What they Do 
-a   | List all active tasks
-c   | List All Completed Tasks
-d {}| {} represents a date. First put the -d and then put in a date. Framework will spit out all completed tasks after that date
-----------------------------------------
Help Command
task --help 

Clarificaiton questions

    1. What is a task? What information should a task store?
    ANSWER
    ======================================================
    Feature      | Optional/required | Default   | Options
    a status     | Optional          | Created   | Created/Completed/Rejected
    creation date| Optional          | time.now()
    a category   | Optional          | Primary   | Primary or user provided string
    task_type    | Optional          | Parent    | Either parent or subtask
    a title      | Required          | None      
    Completed Day| Optional          | time when task was updated
    Descrption   | Optional          | ""
    a deadline   | Optional          | 100 years away


    2. Where should we persist the information of these tasks? 
    ANSWER
    =====================================================
    A database seems to make sense here as we have a clear model for each task and will
    make authentication in the near future easier. Given that we have a category-view and plans for subtasks,
    we should use a relational structure to represent this data as we have a lot of relationships in our data

    3. How do we want to represent tasks in a cli?
    Standard STDOUT should be fine for viewing of information
    Writing should be handled via an i

    4. What open source frameworks/libraries make sense for this project?
     What should the open source frameworks be responsible for? 
     ANSWER
     ==================================================================================
     I am thinking we can just use a libary for styling and handle the rest ourselves for learning purposes.


```
### What all do I need to learn to accomplish this task?
- Golang STDIN 
- Golang STDOUT
- Using Postgres with golang
- Golang Interfaces
- Golang Enums


# Stretch Goals
- Sub tasks
- Query Tasks by date
- Remind user of pending tasks by a time
- user authentication
