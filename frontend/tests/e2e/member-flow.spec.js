import { test, expect } from '@playwright/test';

test.describe('Member Management', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('http://localhost:5173');
  });

  test('should display initial members from seed data', async ({ page }) => {
    const memberItems = page.locator('.member-item');
    await expect(memberItems).not.toHaveCount(0); // Wait for seed data
  });

  test('should add a new member', async ({ page }) => {
    await page.fill('input[placeholder="Name"]', 'Bob Tester');
    await page.fill('input[placeholder="Email"]', 'bob@test.com');
    await page.click('button:has-text("Add Member")');

    const memberItems = page.locator('.member-item');
    await expect(memberItems).toHaveCount(2);
    await expect(page.locator('.member-item:last-child')).toContainText('Bob Tester');
  });

  test('should delete a member', async ({ page }) => {
    const initialCount = await page.locator('.member-item').count();
    await page.click('.member-item:first-child button:has-text("Delete")');
    
    // Wait for the count to decrease
    const memberItems = page.locator('.member-item');
    await expect(memberItems).toHaveCount(initialCount - 1);
  });
});
