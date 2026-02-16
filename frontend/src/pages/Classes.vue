<template>
  <div class="card">
    <h1>Gym Classes</h1>
    
    <div v-if="loading">Loading...</div>
    <div v-else class="classes-list">
      <div v-for="c in classes" :key="c.id" class="class-item">
        <h3>{{ c.name }}</h3>
        <p>Trainer: {{ c.trainer }}</p>
        <p>Schedule: {{ c.schedule }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { fetchClasses } from '../services/graphql';

const classes = ref([]);
const loading = ref(true);

onMounted(async () => {
  try {
    classes.value = await fetchClasses();
  } finally {
    loading.value = false;
  }
});
</script>

<style scoped>
.classes-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}
.class-item {
  padding: 1rem;
  background: rgba(255,255,255,0.05);
  border-radius: 8px;
  text-align: left;
}
</style>
