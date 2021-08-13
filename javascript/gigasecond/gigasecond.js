const gigaSecondInMs = 10 ** 9 * 1000;

export const gigasecond = (time) => {
  const unixEpochInMs = time.getTime();
  return new Date(unixEpochInMs + gigaSecondInMs);
};
