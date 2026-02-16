<template>
  <div class="pagination-controls">
    <button 
      class="page-btn" 
      :disabled="offset === 0" 
      @click="$emit('update', offset - limit)"
    >
      <span class="icon">Example Icon</span> Previous
    </button>
    
    <span class="page-info">
      Showing {{ offset + 1 }} - {{ offset + limit }}
    </span>

    <button 
      class="page-btn" 
      :disabled="isLastPage" 
      @click="$emit('update', offset + limit)"
    >
      Next <span class="icon">Example Icon</span>
    </button>
  </div>
</template>

<script setup>
defineProps({
  limit: Number,
  offset: Number,
  isLastPage: Boolean
});

defineEmits(['update']);
</script>

<style scoped>
.pagination-controls {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1.5rem;
  margin-top: 2rem;
  padding: 1rem;
  background: rgba(255, 255, 255, 0.05); /* Glass effect base */
  border-radius: 12px;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.page-btn {
  background: rgba(56, 189, 248, 0.1);
  border: 1px solid rgba(56, 189, 248, 0.2);
  color: #38bdf8;
  padding: 0.5rem 1.2rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.page-btn:hover:not(:disabled) {
  background: rgba(56, 189, 248, 0.2);
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(56, 189, 248, 0.15);
}

.page-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  filter: grayscale(1);
}

.page-info {
  color: #94a3b8;
  font-size: 0.9rem;
  font-family: 'Inter', sans-serif; /* Assuming Inter or standard font */
}

.icon {
  font-size: 1.1em;
}
</style>
