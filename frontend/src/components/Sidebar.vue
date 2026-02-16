<template>
  <div class="sidebar">
    <h4>Recent Plans</h4>
    <ul>
      <li v-for="p in smallPlans" :key="p.id">
        {{ p.name }} - ${{ p.total_price }}
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

const smallPlans = ref([]);

onMounted(async () => {
  const res = await axios.get('http://localhost:8080/api/plans/getall');
  smallPlans.value = res.data.slice(0, 2);
});
</script>

<style scoped>
.sidebar {
  padding: 1rem;
  background: rgba(0,0,0,0.2);
  border-radius: 8px;
  text-align: left;
}
ul { list-style: none; padding: 0; }
li { font-size: 0.9rem; margin-bottom: 0.5rem; }
</style>
