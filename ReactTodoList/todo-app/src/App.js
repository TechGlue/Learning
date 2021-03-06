import React from "react";
import Todo from "./components/Todo";
import FilterButton from "./components/FilterButtons";
import Form from "./components/Form";

function App(props) {
  function addTask(name) {
    alert(name);
  }

  //for looping through our array of tasks so we can have them available.
  const taskList = props.tasks?.map((task) => (
    //keys are managed by react simliar to an ID
    <Todo
      id={task.id}
      name={task.name}
      completed={task.completed}
      key={task.id}
    />
  ));

  return (
    <div className="todoapp stack-large">
      <h1>TodoMatic</h1>
      <Form addTask={addTask} />
      <div className="filters btn-group stack-exception">
        <FilterButton />
        <FilterButton />
        <FilterButton />
      </div>
      <h2 id="list-heading">3 tasks remaining</h2>
      <ul
        role="list"
        className="todo-list stack-large stack-exception"
        aria-labelledby="list-heading"
      >
        {taskList}
      </ul>
    </div>
  );
}

export default App;
