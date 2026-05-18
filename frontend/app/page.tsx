'use client'

import { useState, useEffect } from 'react'
import { todoService, type Todo } from './services/todoService'
import styles from './page.module.css'

export default function Home() {
  const [todos, setTodos] = useState<Todo[]>([])
  const [title, setTitle] = useState<string>('')
  const [loading, setLoading] = useState<boolean>(true)

  useEffect(() => {
    loadTodos()
  }, [])

  async function loadTodos(): Promise<void> {
    try {
      setLoading(true)
      const data = await todoService.getAll()
      setTodos(data)
    } catch (err) {
      console.error('Failed to load todos')
    } finally {
      setLoading(false)
    }
  }

  async function addTodo(e: React.FormEvent<HTMLFormElement>): Promise<void> {
    e.preventDefault()
    if (!title.trim()) return

    try {
      await todoService.create(title)
      setTitle('')
      loadTodos()
    } catch (err) {
      console.error('Failed to add todo')
    }
  }

  async function toggleTodo(id: number, completed: boolean): Promise<void> {
    try {
      await todoService.update(id, !completed)
      loadTodos()
    } catch (err) {
      console.error('Failed to toggle todo')
    }
  }

  async function deleteTodo(id: number): Promise<void> {
    try {
      await todoService.delete(id)
      loadTodos()
    } catch (err) {
      console.error('Failed to delete todo')
    }
  }

  return (
    <main className={styles.container}>
      <h1>Go Todo</h1>
      
      <form onSubmit={addTodo} className={styles.form}>
        <input
          type="text"
          placeholder="New todo"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />
        <button type="submit">Add</button>
      </form>

      {loading ? (
        <p>Loading...</p>
      ) : (
        <ul className={styles.todoList}>
          {todos.length === 0 ? (
            <li className={styles.empty}>No todos yet. Create one!</li>
          ) : (
            todos.map((todo) => (
              <li key={todo.id} className={todo.completed ? styles.completed : ''}>
                <input
                  type="checkbox"
                  checked={todo.completed}
                  onChange={() => toggleTodo(todo.id, todo.completed)}
                />
                <span>{todo.title}</span>
                <button
                  className={styles.deleteBtn}
                  onClick={() => deleteTodo(todo.id)}
                >
                  Delete
                </button>
              </li>
            ))
          )}
        </ul>
      )}
    </main>
  )
}
