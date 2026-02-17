<template>
  <div class="member-profile">
    <div class="header">
      <h1>Member Profile</h1>
      <router-link to="/" class="btn-secondary">Back to List</router-link>
    </div>

    <div v-if="loading" class="loading">Loading...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <div v-else class="content">
      <!-- Info Card -->
      <div class="glass-panel info-card">
        <div class="avatar-placeholder">{{ member.name ? member.name.charAt(0) : '?' }}</div>
        <div class="details">
          <h2>{{ member.name }}</h2>
          <p class="email">{{ member.email }}</p>
          <div class="meta">
            <span class="badge" :class="member.status">{{ member.status }}</span>
            <span class="date">Joined: {{ formatDate(member.joined_at) }}</span>
          </div>
        </div>
      </div>

      <!-- Invoices Section -->
      <div class="glass-panel invoices-section">
        <h3>Invoice History</h3>
        
        <table class="glass-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>Date</th>
              <th>Amount</th>
              <th>Status</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="inv in invoices" :key="inv.id">
              <td>#{{ inv.id }}</td>
              <td>{{ formatDate(inv.date) }}</td>
              <td>${{ inv.amount }}</td>
              <td>
                <span class="status-chip" :class="inv.status">{{ inv.status }}</span>
              </td>
            </tr>
            <tr v-if="invoices.length === 0">
              <td colspan="4" class="text-center">No invoices found.</td>
            </tr>
          </tbody>
        </table>

        <Pagination 
          v-if="invoices.length > 0 || offset > 0"
          :limit="limit" 
          :offset="offset" 
          :isLastPage="isLastPage" 
          @update="handlePageUpdate" 
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getMember, getMemberInvoices } from '../services/api';
import Pagination from '../components/Pagination.vue';

const props = defineProps(['id']);
const member = ref(null);
const invoices = ref([]);
const loading = ref(true);
const error = ref(null);

const limit = ref(5);
const offset = ref(0);
const isLastPage = ref(false);

const loadData = async () => {
  loading.value = true;
  try {
    const mRes = await getMember(props.id);
    member.value = mRes.data;

    await loadInvoices();
  } catch (e) {
    error.value = "Failed to load member profile";
    console.error(e);
  } finally {
    loading.value = false;
  }
};

const loadInvoices = async () => {
  try {
    const iRes = await getMemberInvoices(props.id, limit.value, offset.value);
    invoices.value = iRes.data || [];
    isLastPage.value = (iRes.data || []).length < limit.value;
  } catch (e) {
    console.error(e);
  }
};

const handlePageUpdate = (newOffset) => {
  offset.value = newOffset;
  loadInvoices();
};

const formatDate = (dateStr) => {
  if (!dateStr) return 'N/A';
  return new Date(dateStr).toLocaleDateString();
};

onMounted(loadData);
</script>

<style scoped>
.member-profile {
  max-width: 900px;
  margin: 0 auto;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.btn-secondary {
  background: rgba(255, 255, 255, 0.1);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 8px;
  text-decoration: none;
  transition: background 0.3s;
}
.btn-secondary:hover {
  background: rgba(255, 255, 255, 0.2);
}

.info-card {
  display: flex;
  align-items: center;
  gap: 2rem;
  padding: 2.5rem;
  margin-bottom: 2rem;
}

.avatar-placeholder {
  width: 100px;
  height: 100px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 3rem;
  font-weight: bold;
  color: white;
  box-shadow: 0 4px 15px rgba(0,0,0,0.3);
}

.details h2 {
  margin: 0 0 0.5rem 0;
  font-size: 2rem;
}

.email {
  color: var(--text-secondary);
  margin-bottom: 1.5rem;
  font-size: 1.1rem;
}

.meta {
  display: flex;
  gap: 1.5rem;
  align-items: center;
}

.badge {
  padding: 6px 16px;
  border-radius: 20px;
  font-size: 0.9rem;
  text-transform: uppercase;
  font-weight: bold;
  letter-spacing: 0.5px;
}

.badge.active {
  background: rgba(76, 175, 80, 0.2);
  color: #81c784;
  border: 1px solid rgba(76, 175, 80, 0.3);
}

.invoices-section h3 {
  margin-bottom: 1.5rem;
  padding-bottom: 0.5rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.status-chip {
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 0.85rem;
  text-transform: capitalize;
}
.status-chip.paid { color: #81c784; }
.status-chip.pending { color: #ffb74d; }

.text-center { text-align: center; color: var(--text-secondary); }
</style>
