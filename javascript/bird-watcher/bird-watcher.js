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
const sum = (previous, current) => previous + current;

export function totalBirdCount(birdsPerDay) {
  return birdsPerDay.reduce(sum);
}

/**
 * Calculates the total number of birds seen in a specific week.
 *
 * @param {number[]} birdsPerDay
 * @param {number} week
 * @returns {number} birds counted in the given week
 */
export function birdsInWeek(birdsPerDay, week) {
  let endIndex = week * 7;
  let startIndex = endIndex - 7;
  return birdsPerDay.slice(startIndex, endIndex).reduce(sum);
}

/**
 * Fixes the counting mistake by increasing the bird count
 * by one for every second day.
 *
 * @param {number[]} birdsPerDay
 * @returns {number[]} corrected bird count data
 */
export function fixBirdCountLog(birdsPerDay) {
  birdsPerDay.forEach((day, idx) => {
    birdsPerDay[idx] = idx % 2 == 0 ? day + 1 : day;
  });
  return birdsPerDay;
}
