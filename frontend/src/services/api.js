import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
});

export const fetchMembers = (status, limit = 5, offset = 0) => {
  return api.get(`/members?status=${status || ''}&limit=${limit}&offset=${offset}`);
};

export const addMember = (data) => {
  return api.post('/members', data);
};

export const deleteUser = (id) => {
  return api.delete(`/members/${id}`);
};

export const getMember = (id) => {
  return api.get(`/members/${id}`);
};

export const getMemberInvoices = (id, limit = 5, offset = 0) => {
  return api.get(`/members/${id}/invoices?limit=${limit}&offset=${offset}`);
};

export const get_plans = () => {
  return api.get('/plans/getall');
};

export const create_plan = (data) => {
  return api.post('/plans_create', data);
};

export const fetchInvoices = (limit = 5, offset = 0) => {
  return api.get(`/invoices?limit=${limit}&offset=${offset}`);
};

export const createInvoice = (data) => {
  return api.post('/invoices/create', data);
};

export const payInvoice = (id) => {
  return api.post(`/invoices/pay/${id}`);
};

export const getDashboardData = () => {
  return api.get('/dashboard');
};
