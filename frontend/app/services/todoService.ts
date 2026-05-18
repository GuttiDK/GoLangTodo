export interface Todo {
  id: number
  title: string
  completed: boolean
}

const API_BASE = '/api/todos'

export const todoService = {
  async getAll(): Promise<Todo[]> {
    try {
      const res = await fetch(API_BASE)
      if (!res.ok) throw new Error('Failed to fetch todos')
      const data = await res.json()
      return data || []
    } catch (err) {
      console.error('Error loading todos:', err)
      throw err
    }
  },

  async create(title: string): Promise<Todo> {
    try {
      const res = await fetch(API_BASE, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title: title.trim() })
      })
      if (!res.ok) throw new Error('Failed to create todo')
      return await res.json()
    } catch (err) {
      console.error('Error adding todo:', err)
      throw err
    }
  },

  async update(id: number, completed: boolean): Promise<Todo> {
    try {
      const res = await fetch(`${API_BASE}/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ completed })
      })
      if (!res.ok) throw new Error('Failed to update todo')
      return await res.json()
    } catch (err) {
      console.error('Error updating todo:', err)
      throw err
    }
  },

  async delete(id: number): Promise<void> {
    try {
      const res = await fetch(`${API_BASE}/${id}`, { method: 'DELETE' })
      if (!res.ok) throw new Error('Failed to delete todo')
    } catch (err) {
      console.error('Error deleting todo:', err)
      throw err
    }
  }
}
