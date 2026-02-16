<template>
  <div class="billing-page">
    <div class="glass card">
      <div class="header">
        <h1>Billing & Invoices</h1>
        <button @click="handleCreate" class="btn-create">
          Create Invoice
        </button>
      </div>

      <div class="list-container">
        <transition-group name="list" tag="div" class="invoice-list">
          <div v-for="inv in invoices" :key="inv.id" class="invoice-item glass-panel">
            <div class="invoice-details">
              <span class="member-name">{{ inv.member_name }}</span>
              <span class="amount">${{ inv.amount.toFixed(2) }}</span>
              <span class="date">{{ formatDate(inv.date) }}</span>
            </div>
            
            <div class="actions">
              <span :class="['status-badge', inv.status]">
                {{ inv.status }}
              </span>
              <button 
                v-if="inv.status === 'pending'"
                @click="handlePay(inv.id)"
                class="btn-pay"
              >
                Mark Paid
              </button>
            </div>
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
import { fetchInvoices, payInvoice, createInvoice } from '../services/api';
import Pagination from '../components/Pagination.vue';

const invoices = ref([]);
const limit = ref(5);
const offset = ref(0);
const isLastPage = ref(false);

const loadInvoices = async () => {
  try {
    const res = await fetchInvoices(limit.value, offset.value);
    invoices.value = res.data || [];
    isLastPage.value = (res.data || []).length < limit.value;
  } catch (e) {
    console.error("Failed to load invoices");
  }
};

onMounted(loadInvoices);

watch(offset, loadInvoices);

const handlePagination = (newOffset) => {
  if (newOffset < 0) return;
  offset.value = newOffset;
};

const handlePay = async (id) => {
  try {
    await payInvoice(id);
    // Optimistic update
    const inv = invoices.value.find(i => i.id === id);
    if (inv) inv.status = 'paid';
  } catch (e) {
    alert("Payment failed");
  }
};

const handleCreate = async () => {
  alert("Create Invoice feature coming soon!");
};

const formatDate = (dateStr) => {
  if (!dateStr || dateStr === "0001-01-01T00:00:00Z") return "N/A";
  return new Date(dateStr).toLocaleDateString();
};
</script>

<style scoped>
.billing-page {
  max-width: 800px;
  margin: 0 auto;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.btn-create {
  background: rgba(56, 189, 248, 0.1);
  border: 1px solid rgba(56, 189, 248, 0.2);
  color: #38bdf8;
  padding: 0.5rem 1rem;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-create:hover {
  background: rgba(56, 189, 248, 0.2);
}

.invoice-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  margin-bottom: 0.75rem;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.05);
  border-radius: 12px;
}

.invoice-details {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.member-name {
  font-weight: 600;
  color: #f1f5f9;
}

.amount {
  font-size: 1.1rem;
  color: #38bdf8;
  font-weight: 700;
}

.date {
  font-size: 0.8rem;
  color: #94a3b8;
}

.actions {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.status-badge {
  padding: 0.25rem 0.75rem;
  border-radius: 999px;
  font-size: 0.75rem;
  font-weight: 600;
  text-transform: uppercase;
}

.status-badge.paid {
  background: rgba(34, 197, 94, 0.2);
  color: #4ade80;
}

.status-badge.pending {
  background: rgba(251, 191, 36, 0.2);
  color: #fbbf24;
}

.btn-pay {
  background: linear-gradient(135deg, #22c55e 0%, #16a34a 100%);
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  color: white;
  font-weight: 600;
  cursor: pointer;
  transition: opacity 0.2s;
}

.btn-pay:hover {
  opacity: 0.9;
}
</style>

