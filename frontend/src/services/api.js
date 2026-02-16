import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
});

export const fetchMembers = (status) => {
  return api.get(`/members?status=${status || ''}`);
};

export const addMember = (data) => {
  return api.post('/members', data);
};

export const deleteUser = (id) => {
  return api.delete(`/members/${id}`);
};

export const get_plans = () => {
  return api.get('/plans/getall');
};

export const create_plan = (data) => {
  return api.post('/plans_create', data);
};

export const fetchInvoices = () => {
  return api.get('/invoices');
};

export const createInvoice = (data) => {
  return api.post('/invoices/create', data);
};

export const payInvoice = (id) => {
  return api.post(`/invoices/pay/${id}`);
};
