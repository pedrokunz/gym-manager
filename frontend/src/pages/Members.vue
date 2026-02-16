<template>
  <div class="members-page">
    <div class="glass card">
      <div class="header">
        <h1>Member Management</h1>
        <span class="badge">{{ members.length }} Active</span>
      </div>
      
      <div class="form-container">
        <div class="input-group">
          <input v-model="newMember.name" placeholder="Full Name" class="glass-input" />
          <input v-model="newMember.email" placeholder="Email Address" class="glass-input" />
        </div>
        <button @click="handleAddMember" class="btn-primary">
          Add Member
        </button>
      </div>
      <p v-if="error" class="error-msg">{{ error }}</p>

      <div class="list-container">
        <transition-group name="list" tag="div" class="member-list">
          <div v-for="m in members" :key="m.id" class="member-item glass-panel">
<<<<<<< HEAD
            <router-link :to="`/members/${m.id}`" class="member-info">
              <span class="member-name">{{ m.name }}</span>
              <span class="member-email">{{ m.email }}</span>
            </router-link>
=======
            <div class="member-info">
              <span class="member-name">{{ m.name }}</span>
              <span class="member-email">{{ m.email }}</span>
            </div>
>>>>>>> 8804b81 (feat(frontend): add pagination, glassmorphism and animations)
            <button @click="handleDelete(m.id)" class="btn-danger">
              Remove
            </button>
          </div>
        </transition-group>
      </div>

      <Pagination 
        :limit="limit" 
        :offset="offset" 
        :isLastPage="isLastPage" 
        @update="handlePagination" 
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { fetchMembers, addMember, deleteUser } from '../services/api';
import Pagination from '../components/Pagination.vue';

const members = ref([]);
const newMember = ref({ name: '', email: '' });
const error = ref('');
const limit = ref(5);
const offset = ref(0);
const isLastPage = ref(false);

const loadMembers = async () => {
  try {
    const res = await fetchMembers('', limit.value, offset.value);
    members.value = res.data || [];
    // Simple heuristic for last page: if result count < limit.
    // Ideally backend returns total count.
    isLastPage.value = (res.data || []).length < limit.value;
  } catch (e) {
    error.value = "Failed to load members";
  }
};

onMounted(loadMembers);

watch(offset, loadMembers);

const handlePagination = (newOffset) => {
  if (newOffset < 0) return;
  offset.value = newOffset;
};

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
    // Refresh list or add locally if on first page?
    // Let's refresh to be safe and consistent with pagination.
    loadMembers(); 
    newMember.value = { name: '', email: '' };
    error.value = '';
  } catch (e) {
    error.value = "Failed to add member";
  }
};

const handleDelete = async (id) => {
  try {
    await deleteUser(id);
    loadMembers(); // Refresh to keep pagination correct
  } catch (e) {
    error.value = "Failed to delete member";
  }
};
</script>

<style scoped>
<<<<<<< HEAD
/* ... */
.member-info {
  display: flex;
  flex-direction: column;
  text-decoration: none;
  flex: 1;
  cursor: pointer;
}

=======
.members-page {
  max-width: 800px;
  margin: 0 auto;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.badge {
  background: rgba(56, 189, 248, 0.2);
  color: #38bdf8;
  padding: 0.25rem 0.75rem;
  border-radius: 999px;
  font-size: 0.875rem;
  font-weight: 600;
}

.form-container {
  display: flex;
  gap: 1rem;
  margin-bottom: 2rem;
  flex-wrap: wrap;
}

.input-group {
  display: flex;
  gap: 1rem;
  flex: 1;
}

.glass-input {
  flex: 1;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.1);
  padding: 0.75rem 1rem;
  border-radius: 8px;
  color: white;
  transition: all 0.3s ease;
}

.glass-input:focus {
  outline: none;
  border-color: #38bdf8;
  background: rgba(255, 255, 255, 0.1);
}

.member-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  margin-bottom: 0.75rem;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.05);
  border-radius: 12px;
  transition: all 0.3s ease;
}

.member-item:hover {
  background: rgba(255, 255, 255, 0.08);
  transform: translateX(5px);
}

.member-info {
  display: flex;
  flex-direction: column;
}

>>>>>>> 8804b81 (feat(frontend): add pagination, glassmorphism and animations)
.member-name {
  font-weight: 600;
  color: #f1f5f9;
}

.member-email {
  font-size: 0.875rem;
  color: #94a3b8;
}

.btn-primary {
  background: linear-gradient(135deg, #38bdf8 0%, #0ea5e9 100%);
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  color: white;
  font-weight: 600;
  cursor: pointer;
  transition: opacity 0.2s;
}

.btn-primary:hover {
  opacity: 0.9;
}

.btn-danger {
  background: rgba(248, 113, 113, 0.1);
  border: 1px solid rgba(248, 113, 113, 0.2);
  color: #f87171;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-danger:hover {
  background: rgba(248, 113, 113, 0.2);
}

.error-msg {
  color: #f87171;
  margin-bottom: 1rem;
  padding: 0.5rem;
  background: rgba(248, 113, 113, 0.1);
  border-radius: 6px;
}

/* Animations handled by global style.css via transition-group="list" */
/* But we can ensure they are referenced here or global */
</style>

