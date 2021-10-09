<script setup>
import { defineProps, defineEmits } from 'vue'

const emit = defineEmits(['delete-todo'])
defineProps({
  todos: Array
})

const deleteTodo = (id) => {
  emit('delete-todo', id)
}
</script>

<template>
  <div class="row">
    <div class="eight wide column row">
      <div class="column">

        <div class="ui middle aligned divided list">

          <div class="item" v-for="(todo, index) in todos" :key="'item-' + index">
            <div class="right floated content">
              <button class="ui icon button" @click="deleteTodo(todo.id)">
                <i class="trash alternate icon"></i>
              </button>
            </div>

            <div class="content"><i class="large sticky note icon"></i>{{ todo.body }}</div>
          </div>
          
        </div>

      </div>
    </div>
  </div>
</template>
