import { describe, it, expect } from 'vitest'

const calculateDiscountedPrice = (total, months) => {
  if (months >= 12) {
    return (total * 0.9 / months).toFixed(2)
  }
  return (total / months).toFixed(2)
}

describe('Pricing Logic', () => {
  it('calculates monthly price correctly for 1 month', () => {
    expect(calculateDiscountedPrice(50, 1)).toBe('50.00')
  })

  it('applies 10% discount for 12 months', () => {
    expect(calculateDiscountedPrice(480, 12)).toBe('36.00')
  })
})
