<template>
  <div class="dashboard-page">
    <div class="header">
      <h1>Management Dashboard</h1>
      <p class="subtitle">Overview of your gym's performance</p>
    </div>

    <div v-if="loading" class="loading">
      <div class="spinner"></div>
      <p>Crunching the numbers...</p>
    </div>
    
    <div v-else-if="error" class="error-msg">
      {{ error }}
    </div>

    <div v-else class="dashboard-content">
      <!-- Stats Grid -->
      <div class="stats-grid">
        <div class="glass-panel stat-card">
          <div class="stat-icon members">👥</div>
          <div class="stat-info">
            <span class="stat-label">Total Members</span>
            <span class="stat-value">{{ stats.total_members }}</span>
          </div>
        </div>

        <div class="glass-panel stat-card">
          <div class="stat-icon active">✅</div>
          <div class="stat-info">
            <span class="stat-label">Active Members</span>
            <span class="stat-value">{{ stats.active_members }}</span>
          </div>
        </div>

        <div class="glass-panel stat-card">
          <div class="stat-icon revenue">💰</div>
          <div class="stat-info">
            <span class="stat-label">Total Revenue</span>
            <span class="stat-value">${{ stats.total_revenue.toLocaleString() }}</span>
          </div>
        </div>

        <div class="glass-panel stat-card">
          <div class="stat-icon monthly">📅</div>
          <div class="stat-info">
            <span class="stat-label">Monthly Goal</span>
            <span class="stat-value">85%</span>
          </div>
        </div>
      </div>

      <div class="main-sections">
        <!-- Recent Invoices -->
        <div class="glass-panel activity-section">
          <h3>Recent Payments</h3>
          <div class="invoice-list">
            <div v-for="inv in stats.recent_invoices" :key="inv.id" class="invoice-item">
              <div class="inv-main">
                <span class="inv-member">{{ inv.member_name }}</span>
                <span class="inv-date">{{ formatDate(inv.date) }}</span>
              </div>
              <div class="inv-amount-status">
                <span class="inv-amount">${{ inv.amount }}</span>
                <span class="status-pill" :class="inv.status">{{ inv.status }}</span>
              </div>
            </div>
            <div v-if="!stats.recent_invoices?.length" class="empty-state">
              No recent activity found.
            </div>
          </div>
        </div>

        <!-- Quick Actions or Chart Placeholder -->
        <div class="glass-panel charts-section">
          <h3>Member Growth</h3>
          <div class="chart-placeholder">
            <div class="bar" style="height: 40%"></div>
            <div class="bar" style="height: 60%"></div>
            <div class="bar" style="height: 50%"></div>
            <div class="bar" style="height: 80%"></div>
            <div class="bar" style="height: 70%"></div>
            <div class="bar" style="height: 90%"></div>
          </div>
          <p class="chart-label">Last 6 Months Engagement</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getDashboardData } from '../services/api';

const stats = ref({
  total_members: 0,
  active_members: 0,
  total_revenue: 0,
  recent_invoices: []
});
const loading = ref(true);
const error = ref(null);

const loadDashboard = async () => {
  try {
    const res = await getDashboardData();
    stats.value = res.data;
  } catch (e) {
    error.value = "Failed to load dashboard data.";
    console.error(e);
  } finally {
    loading.value = false;
  }
};

const formatDate = (dateStr) => {
  if (!dateStr) return 'N/A';
  return new Date(dateStr).toLocaleDateString();
};

onMounted(loadDashboard);
</script>

<style scoped>
.dashboard-page {
  padding: 1rem;
  animation: fadeIn 0.5s ease-out;
}

.header {
  margin-bottom: 2.5rem;
}

.header h1 {
  font-size: 2.5rem;
  background: linear-gradient(to right, #fff, #94a3b8);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  margin-bottom: 0.5rem;
}

.subtitle {
  color: #94a3b8;
  font-size: 1.1rem;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2.5rem;
}

.stat-card {
  display: flex;
  align-items: center;
  gap: 1.25rem;
  padding: 1.5rem;
  transition: transform 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-5px);
}

.stat-icon {
  font-size: 2rem;
  width: 50px;
  height: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 12px;
}

.stat-info {
  display: flex;
  flex-direction: column;
}

.stat-label {
  color: #94a3b8;
  font-size: 0.9rem;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.stat-value {
  font-size: 1.5rem;
  font-weight: 700;
  color: #fff;
}

.main-sections {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
}

@media (max-width: 768px) {
  .main-sections {
    grid-template-columns: 1fr;
  }
}

.activity-section, .charts-section {
  padding: 1.5rem;
}

h3 {
  margin-bottom: 1.5rem;
  color: #f1f5f9;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  padding-bottom: 0.75rem;
}

.invoice-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.invoice-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.75rem;
  background: rgba(255, 255, 255, 0.02);
  border-radius: 8px;
  border: 1px solid rgba(255, 255, 255, 0.05);
}

.inv-main {
  display: flex;
  flex-direction: column;
}

.inv-member {
  font-weight: 600;
}

.inv-date {
  font-size: 0.8rem;
  color: #64748b;
}

.inv-amount-status {
  text-align: right;
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.inv-amount {
  font-weight: 700;
  color: #38bdf8;
}

.status-pill {
  font-size: 0.7rem;
  padding: 2px 8px;
  border-radius: 99px;
  text-transform: uppercase;
}

.status-pill.paid { background: rgba(34, 197, 94, 0.2); color: #4ade80; }
.status-pill.pending { background: rgba(234, 179, 8, 0.2); color: #fde047; }

.chart-placeholder {
  height: 200px;
  display: flex;
  align-items: flex-end;
  gap: 1rem;
  padding: 1rem;
}

.bar {
  flex: 1;
  background: linear-gradient(to top, #38bdf8, #818cf8);
  border-radius: 4px 4px 0 0;
  opacity: 0.7;
}

.chart-label {
  text-align: center;
  color: #64748b;
  font-size: 0.9rem;
  margin-top: 1rem;
}

.loading {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 300px;
  color: #94a3b8;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 4px solid rgba(255, 255, 255, 0.1);
  border-top: 4px solid #38bdf8;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 1rem;
}

@keyframes spin { 0% { transform: rotate(0deg); } 100% { transform: rotate(360deg); } }
@keyframes fadeIn { from { opacity: 0; transform: translateY(10px); } to { opacity: 1; transform: translateY(0); } }
</style>
