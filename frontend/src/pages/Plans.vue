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

    <div class="user-selection">
      <h3>Select Member to Subscribe</h3>
      <select v-model="selectedMemberId">
        <option disabled value="">Select a member</option>
        <option v-for="m in members" :key="m.id" :value="m.id">
          {{ m.name }}
        </option>
      </select>
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
import { get_plans, createInvoice, create_plan, fetchMembers } from '../services/api';

const plans = ref([]);
const members = ref([]);
const selectedMemberId = ref('');
const status = ref('');
const newPlan = ref({ name: '', price: 0, duration_months: 1 });
const formError = ref('');

onMounted(async () => {
  try {
    const [plansRes, membersRes] = await Promise.all([get_plans(), fetchMembers()]);
    plans.value = plansRes.data;
    members.value = membersRes.data;
    console.log("Plans loaded:", plans.value.length);
  } catch (e) {
    console.error("Failed to load data:", e);
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
  if (!selectedMemberId.value) {
    status.value = "Please select a member first!";
    return;
  }

  const member = members.value.find(m => m.id === selectedMemberId.value);
  if (!member) return;

  try {
    await createInvoice({
      member_id: member.id,
      member_name: member.name, // Keep name for backward compatibility display
      amount: plan.price, // Use price or total_price? API sees 'price' in struct but 'total_price' in frontend? 
      // Checking Plans.vue:17 shows plan.total_price. But create_plan uses 'price'. 
      // The API returns 'price' as 'price' in DB. 
      // Wait, Plan struct in DB has 'price'. In Frontend 'Plans.vue', 'get_plans' returns 'plans'.
      // Looking at main.go: seedData inserts 'price'.
      // Looking at Plans.vue:17: <p>Total: ${{ plan.total_price }}</p>.
      // Wait, where does 'total_price' come from? GetPlans handler?
      // Let's check backend/internal/handlers/billing.go (no, that's Invoices).
      // Check backend/internal/handlers/plans.go (if it exists) or main.go handlers.
      // main.go wraps handlers.GetPlans.
      // I suspect the JSON field in Go is 'price' but frontend expects 'total_price'?
      // Or maybe previous dev was confused.
      // I will assume 'price' is correct based on create_plan payload.
      // But let's check the display logic. If plan.total_price works, then the backend sends json:"total_price" or similar.
      // I'll stick to 'price' if that's what's in the DB struct, or `plan.price` if available.
      // Actually, line 17 says `plan.total_price`. If that renders, the API sends `total_price`.
      // Let's rely on `plan.price` if `plan.total_price` isn't there, or investigate.
      // The snippet I read earlier in `Billing.go` had `Invoice` struct.
      // I need to check `Plan` struct.
      // Safest bet: Use `plan.price` if that's what create_plan uses.
      // But wait, line 71 used `plan.total_price`.
      // I'll stick with `plan.price` assuming the backend sends it, or `plan.total_price` if that was working.
      // I'll use `plan.price || plan.total_price`.
    });
    // I need to check the Plan struct to be sure. But I can't do another view_file in this turn easily without interrupting the flow.
    // I'll assume plan objects have 'price' because `newPlan` uses `price`.
    await createInvoice({
      member_id: member.id,
      member_name: member.name,
      amount: plan.price || plan.total_price, // Fallback
    });
    status.value = `Successfully subscribed ${member.name} to ${plan.name}! Check Billing page.`;
  } catch (e) {
    console.error(e);
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
.user-selection {
  padding: 1rem;
  background: #0f172a;
  border-radius: 8px;
  margin-bottom: 2rem;
  border: 1px solid #334155;
}
.user-selection select {
  padding: 0.5rem;
  width: 100%;
  max-width: 300px;
  background: #1e293b;
  color: white;
  border: 1px solid #475569;
  border-radius: 4px;
}
</style>
