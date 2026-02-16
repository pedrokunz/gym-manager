<template>
  <div class="layout">
    <nav class="nav">
      <router-link to="/">Members</router-link>
      <router-link to="/plans">Plans</router-link>
      <router-link to="/classes">Classes</router-link>
      <router-link to="/billing">Billing</router-link>
    </nav>
    
    <div class="main-container">
      <aside class="sidebar-wrapper">
        <Sidebar />
      </aside>
      <main>
        <router-view></router-view>
      </main>
    </div>

    <footer>
      &copy; 2026 Gym Manager Portal
    </footer>
  </div>
</template>

<script setup>
import Sidebar from './components/Sidebar.vue';
</script>

<style scoped>
.layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.main-container {
  display: flex;
  gap: 2rem;
  text-align: left;
}

.sidebar-wrapper {
  flex: 0 0 250px;
}

main {
  flex: 1;
}
</style>
