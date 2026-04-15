import { useState } from "react";

const BASE_URL = "https://studious-space-guide-pj496p4j57wjh64gr-8080.app.github.dev"; // 🔥 replace with your backend URL

function App() {
  const [projects, setProjects] = useState([]);
  const [name, setName] = useState("");
  const [tasks, setTasks] = useState([]);
  const [task, setTask] = useState("");
  const [pid, setPid] = useState(null);

  // ✅ Create Project
  const createProject = async () => {
    await fetch(`${BASE_URL}/projects`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name }),
    });

    setName("");
    loadProjects();
  };

  // ✅ Load Projects
  const loadProjects = async () => {
    const res = await fetch(`${BASE_URL}/projects`);
    const data = await res.json();
    setProjects(data.projects || []);
  };

  // ✅ Load Tasks
  const loadTasks = async (id) => {
    setPid(id);
    const res = await fetch(`${BASE_URL}/projects/${id}/tasks`);
    const data = await res.json();
    setTasks(data.tasks || []);
  };

  // ✅ Add Task
  const addTask = async () => {
    await fetch(`${BASE_URL}/projects/${pid}/tasks`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title: task }),
    });

    setTask("");
    loadTasks(pid);
  };

  // ✅ Delete Project
  const deleteProject = async (id) => {
    await fetch(`${BASE_URL}/projects/${id}`, {
      method: "DELETE",
    });

    loadProjects();
  };

  // ✅ Delete Task
  const deleteTask = async (id) => {
    await fetch(`${BASE_URL}/tasks/${id}`, {
      method: "DELETE",
    });

    loadTasks(pid);
  };

  return (
    <div style={{ padding: 20 }}>
      <h2>🚀 TaskFlow</h2>

      {/* Create Project */}
      <h3>Create Project</h3>
      <input
        placeholder="Project name"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <button onClick={createProject}>Create</button>

      <br /><br />

      {/* Load Projects */}
      <button onClick={loadProjects}>Load Projects</button>

      <h3>Projects</h3>
      {projects.map((p) => (
        <div key={p.id} style={{ marginBottom: "10px" }}>
          <b>{p.name}</b>

          <br />

          <button onClick={() => loadTasks(p.id)}>
            View Tasks
          </button>

          <button onClick={() => deleteProject(p.id)}>
            Delete ❌
          </button>
        </div>
      ))}

      <hr />

      {/* Tasks */}
      <h3>Tasks</h3>

      {tasks.map((t) => (
        <div key={t.id} style={{ marginBottom: "10px" }}>
          {t.title} - {t.status}

          <button onClick={() => deleteTask(t.id)}>
            Delete ❌
          </button>
        </div>
      ))}

      {/* Add Task */}
      {pid && (
        <>
          <input
            placeholder="New task"
            value={task}
            onChange={(e) => setTask(e.target.value)}
          />
          <button onClick={addTask}>Add Task</button>
        </>
      )}
    </div>
  );
}

export default App;