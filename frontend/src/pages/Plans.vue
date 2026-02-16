<template>
  <div class="card">
    <h1>Membership Plans</h1>

    <div class="create-form">
      <h3>Add New Plan</h3>
      <input v-model="newPlan.name" placeholder="Plan Name" />
      <input v-model="newPlan.price" type="number" placeholder="Total Price" />
      <input v-model="newPlan.duration_months" type="number" placeholder="Months" />
      <button @click="handleCreatePlan">Save Plan</button>
      <p v-if="formError" class="error">{{ formError }}</p>
    </div>
    
    <div class="plans-grid">
      <div v-for="plan in plans" :key="plan.id" class="plan-card">
        <h3>{{ plan.name }}</h3>
        <p>Total: ${{ plan.total_price }}</p>
        <p>Months: {{ plan.months }}</p>
        
        <div class="pricing-logic">
          <p v-if="plan.months >= 12">
            Discounted Monthly: ${{ (plan.total_price * 0.9 / plan.months).toFixed(2) }}
          </p>
          <p v-else>
            Monthly: ${{ (plan.total_price / plan.months).toFixed(2) }}
          </p>
        </div>
        <button class="buy-btn" @click="subscribe(plan)">Subscribe Now</button>
      </div>
    </div>
    <p v-if="status" class="status">{{ status }}</p>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { get_plans, createInvoice, create_plan } from '../services/api';

const plans = ref([]);
const status = ref('');
const newPlan = ref({ name: '', price: 0, duration_months: 1 });
const formError = ref('');

onMounted(async () => {
  try {
    const res = await get_plans();
    plans.value = res.data;
    console.log("Plans loaded:", plans.value.length);
  } catch (e) {
    console.error("Failed to load plans:", e);
  }
});

const handleCreatePlan = async () => {
  if (newPlan.value.name.length < 2) {
    formError.value = "Short name!";
    return;
  }
  
  await create_plan(newPlan.value);
  const res = await get_plans();
  plans.value = res.data;
  newPlan.value = { name: '', price: 0, duration_months: 1 };
  formError.value = '';
};

const subscribe = async (plan) => {
  try {
    await createInvoice({
      member_name: 'Current User', // Mocked user
      amount: plan.total_price,
    });
    status.value = `Successfully subscribed to ${plan.name}! Check Billing page.`;
  } catch (e) {
    status.value = "Subscription failed.";
  }
};
</script>

<style scoped>
.create-form {
  padding: 1.5rem;
  background: rgba(255,255,255,0.05);
  border-radius: 12px;
  margin-bottom: 2rem;
  display: flex;
  gap: 0.5rem;
  align-items: center;
}
.create-form input {
  padding: 0.5rem;
  border-radius: 4px;
  border: 1px solid #334155;
  background: #0f172a;
  color: white;
}
.create-form button {
  padding: 0.5rem 1rem;
  background: #3b82f6;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
.error { color: #f87171; font-size: 0.8rem; }
.plans-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 1rem;
}
.plan-card {
  padding: 1.5rem;
  background: #1e293b;
  border-radius: 12px;
  border: 1px solid rgba(255,255,255,0.1);
  display: flex;
  flex-direction: column;
}
.pricing-logic {
  margin-top: 1rem;
  font-weight: bold;
  color: #38bdf8;
  flex-grow: 1;
  margin-bottom: 1rem;
}
.buy-btn {
  background: #10b981;
  color: white;
  border: none;
  padding: 0.5rem;
  border-radius: 6px;
  cursor: pointer;
}
.buy-btn:hover { background: #059669; }
.status {
  margin-top: 2rem;
  color: #10b981;
  font-weight: bold;
}
</style>
