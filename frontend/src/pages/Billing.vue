<template>
  <div class="card">
    <h1>Billing & Invoices</h1>
    <div class="invoice-list">
      <div v-for="inv in invoices" :key="inv.id" class="invoice-item">
        <div class="inv-info">
          <strong>#{{ inv.id }} - {{ inv.member_name }}</strong>
          <p>{{ formatDate(inv.date) }}</p>
        </div>
        <div class="inv-status">
          <span :class="inv.status">{{ inv.status.toUpperCase() }}</span>
          <p>${{ inv.amount }}</p>
        </div>
        <button v-if="inv.status === 'pending'" class="pay-btn" @click="handlePay(inv.id)">Pay Now</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { fetchInvoices, payInvoice } from '../services/api';

const invoices = ref([]);

const loadInvoices = async () => {
  const res = await fetchInvoices();
  invoices.value = res.data;
};

onMounted(loadInvoices);

const handlePay = async (id) => {
  try {
    await payInvoice(id);
    await loadInvoices();
  } catch (e) {
    console.error("Payment failed", e);
  }
};

const formatDate = (dateStr) => {
  if (!dateStr || dateStr === "0001-01-01T00:00:00Z") return "N/A";
  return new Date(dateStr).toLocaleDateString();
};
</script>

<style scoped>
.invoice-list { display: flex; flex-direction: column; gap: 1rem; }
.invoice-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  background: #1e293b;
  border-radius: 8px;
  text-align: left;
}
.inv-info { flex: 1; }
.inv-status { margin-right: 1.5rem; text-align: right; }
.pending { color: #fbbf24; font-weight: bold; }
.paid { color: #10b981; font-weight: bold; }
.pay-btn {
  background: #3b82f6;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  cursor: pointer;
}
.pay-btn:hover { background: #2563eb; }
</style>
