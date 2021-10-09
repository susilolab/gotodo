<script setup>
import { ref, defineEmits } from 'vue'

const emit = defineEmits(['save-todo'])
const body = ref('')
const saveTodo = () => {
  emit('save-todo', body.value)
  body.value = ''
}
</script>

<template>
  <h1 class="ui header">Aplikasi Todo</h1>
  <form class="ui form">
    <div class="field">
      <input type="text" v-model="body" placeholder="Ketikan todo" @keydown.enter.prevent="saveTodo">
    </div>
  </form>
</template>
