import { test, expect } from '@playwright/test';

test.describe('Plan Management', () => {
  test.beforeEach(async ({ page }) => {
    page.on('console', msg => console.log(`[${msg.type()}] BROWSER: ${msg.text()}`));
    await page.goto('http://localhost:5173/plans');
    await expect(page.locator('.create-form h3')).toContainText('Add New Plan');
  });

  test('should create a new membership plan', async ({ page }) => {
    const planCards = page.locator('.plan-card');
    await expect(planCards).not.toHaveCount(0); // Wait for seed data
    const initialCount = await planCards.count();

    await page.fill('input[placeholder="Plan Name"]', 'Summer Special');
    await page.fill('input[placeholder="Total Price"]', '100');
    await page.fill('input[placeholder="Months"]', '3');
    await page.click('button:has-text("Save Plan")');

    await expect(page.locator('.plan-card')).toHaveCount(initialCount + 1);
    await expect(page.locator('.plan-card:last-child h3')).toContainText('Summer Special');
  });

  test('should show error for short plan name', async ({ page }) => {
    await page.fill('input[placeholder="Plan Name"]', 'A');
    await page.click('button:has-text("Save Plan")');
    await expect(page.locator('.error')).toContainText('Short name!');
  });
});
