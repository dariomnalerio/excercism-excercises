// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Calculates the total bird count.
 *
 * @param {number[]} birdsPerDay
 * @returns {number} total bird count
 */
export function totalBirdCount(birdsPerDay) {
  let total = 0;
  for (let idx = 0; idx < birdsPerDay.length; idx++) {
    total += birdsPerDay[idx];
  }
  return total;
}

/**
 * Calculates the total number of birds seen in a specific week.
 *
 * @param {number[]} birdsPerDay
 * @param {number} week
 * @returns {number} birds counted in the given week
 */
export function birdsInWeek(birdsPerDay, week) {
  const start = (week - 1) * 7;
  const end = start + 7;
  const weekBirds = birdsPerDay.slice(start, end);
  return totalBirdCount(weekBirds);
}

/**
 * Fixes the counting mistake by increasing the bird count
 * by one for every second day.
 *
 * @param {number[]} birdsPerDay
 * @returns {number[]} corrected bird count data
 */
export function fixBirdCountLog(birdsPerDay) {
  for (let idx = 0; idx < birdsPerDay.length; idx++) {
    if (idx % 2 === 0) {
      birdsPerDay[idx] += 1;
    }
  }
  return birdsPerDay;
}
