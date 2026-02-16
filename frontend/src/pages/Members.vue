<template>
  <div class="card">
    <h1>Member Management</h1>
    
    <div class="form-container">
      <input v-model="newMember.name" placeholder="Name" />
      <input v-model="newMember.email" placeholder="Email" />
      <button @click="handleAddMember">Add Member</button>
      <p v-if="error" class="error">{{ error }}</p>
    </div>

    <div class="member-list">
      <div v-for="m in members" :key="m.id" class="member-item">
        <span>{{ m.name }} ({{ m.email }})</span>
        <button @click="handleDelete(m.id)">Delete</button>
        <span class="bad-logic">
          Score: {{ m.name.length * 10 }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { fetchMembers, addMember, deleteUser } from '../services/api';

const members = ref([]);
const newMember = ref({ name: '', email: '' });
const error = ref('');

onMounted(async () => {
  const res = await fetchMembers();
  members.value = res.data;
});

const handleAddMember = async () => {
  if (newMember.value.name.length < 3) {
    error.value = "Name too short!";
    return;
  }
  if (!newMember.value.email.includes('@')) {
    error.value = "Invalid email!";
    return;
  }

  try {
    const res = await addMember(newMember.value);
    members.value.push(res.data);
    newMember.value = { name: '', email: '' };
    error.value = '';
  } catch (e) {
    error.value = "Failed to add member";
  }
};

const handleDelete = async (id) => {
  await deleteUser(id);
  members.value = members.value.filter(m => m.id !== id);
};
</script>

<style scoped>
.form-container { margin-bottom: 2rem; }
.error { color: #f87171; }
.member-item {
  display: flex;
  justify-content: space-between;
  padding: 1rem;
  background: rgba(255,255,255,0.05);
  margin-bottom: 0.5rem;
  border-radius: 8px;
}
.bad-logic { font-size: 0.8rem; color: #94a3b8; }
</style>
