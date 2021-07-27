export const gigasecond = (time) => {
  let unixEpochInMs = new Date(time) / 1;
  let gigaSecondInMs = 10 ** 9 * 1000;
  return new Date(unixEpochInMs + gigaSecondInMs);
};
