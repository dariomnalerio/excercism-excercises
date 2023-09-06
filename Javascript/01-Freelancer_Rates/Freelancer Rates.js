// @ts-check

/**
 * The day rate, given a rate per hour
 *
 * @param {number} ratePerHour
 * @returns {number} the rate per day
 */

export function dayRate(ratePerHour) {
  return ratePerHour * 8;
}

/**
 * Calculates the number of days in a budget, rounded down
 *
 * @param {number} budget: the total budget
 * @param {number} ratePerHour: the rate per hour
 * @returns {number} the number of days
 */

export function daysInBudget(budget, ratePerHour) {
  const totalHours = budget / ratePerHour;
  const workdays = Math.floor(totalHours / 8);

  return workdays;
}

/**
 * Calculates the discounted rate for large projects, rounded up
 *
 * @param {number} ratePerHour
 * @param {number} numDays: number of days the project spans
 * @param {number} discount: for example 20% written as 0.2
 * @returns {number} the rounded up discounted rate
 */

export function priceWithMonthlyDiscount(ratePerHour, numDays, discount) {
  const fullMonths = Math.floor(numDays / 22);
  const remainingDays = numDays % 22;

  const fullMonthsPrice = fullMonths * 22 * dayRate(ratePerHour);
  const remainingDaysPrice = remainingDays * dayRate(ratePerHour);
  const discountAmount = discount * fullMonthsPrice;

  const finalPrice = Math.ceil(
    fullMonthsPrice - discountAmount + remainingDaysPrice
  );

  return finalPrice;
}
