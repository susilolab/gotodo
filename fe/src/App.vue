<script setup>
import { ref, onMounted } from 'vue'
import TodoApp from './components/TodoApp.vue'
import TodoAppList from './components/TodoAppList.vue'

const todos = ref([])

const loadTodos = () => {
  fetch('http://localhost:8080/todo')
  .then(resp => resp.json())
  .then(resp => todos.value = resp.result)
  .catch(err => console.log)
}

onMounted(() => {
  loadTodos()
})

const deleteTodo = (id) => {
   fetch(`http://localhost:8080/todo/${id}`, {
     method: 'DELETE'
   })
  .then(resp => resp.json())
  .then(resp => {
    if (parseInt(resp.error) === 0) {
      loadTodos()
    }
  })
  .catch(err => console.log)
}

const saveTodo = (text) => {
  fetch('http://localhost:8080/todo', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ body: text, done: false })
  })
  .then(resp => resp.json())
  .then(resp => {
    loadTodos()    
  })
  .catch(err => console.log)
}
</script>

<template>
  <div class="ui grid centered">
    <div class="eight wide column row">
      <div class="column"><TodoApp @save-todo="saveTodo" /></div>
    </div>
    
    <TodoAppList :todos="todos" @delete-todo="deleteTodo" />
  </div>
</template>

<style>
@import '../node_modules/semantic-ui/dist/semantic.min.css';

#app {
  margin-top: 20px;
}
</style>
