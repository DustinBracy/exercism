/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */

export const cookingStatus = (minutes) => {
  switch (true) {
    case minutes == 0:
      return "Lasagna is done.";
    case minutes === undefined:
      return "You forgot to set the timer.";
    default:
      return "Not done, please wait.";
  }
};

export const preparationTime = (layers, layerPrepTime = 2) => {
  return layers.length * layerPrepTime;
};

export const quantities = (layers) => {
  const recipes = { noodles: 50, sauce: 0.2 };
  let requirements = { noodles: 0, sauce: 0 };

  layers.forEach((layer) => {
    if (["noodles", "sauce"].includes(layer)) {
      requirements[layer] += recipes[layer];
    }
  });
  return requirements;
};

export const addSecretIngredient = (friendsList, myList) => {
  myList.push(friendsList[friendsList.length - 1]);
};

export const scaleRecipe = (recipe, portions) => {
  let scaledRecipe = {};
  let defaultPortions = 2;
  Object.keys(recipe).forEach((key) => {
    scaledRecipe[key] = (recipe[key] * portions) / defaultPortions;
  });
  return scaledRecipe;
};
