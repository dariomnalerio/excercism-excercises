/// <reference path="./global.d.ts" />
// @ts-check

export function cookingStatus(time) {
  if (time === 0) {
    return "Lasagna is done.";
  } else if (time === undefined) {
    return "You forgot to set the timer.";
  } else {
    return "Not done, please wait.";
  }
}

export function preparationTime(layers, time = 2) {
  return layers.length * time;
}

export function quantities(layers) {
  const noodles = layers.filter((layer) => layer === "noodles").length * 50;
  const sauce = layers.filter((layer) => layer === "sauce").length * 0.2;
  return { noodles: noodles, sauce: sauce };
}

export function addSecretIngredient(friendList, ownList) {
    ownList.push(friendList[friendList.length - 1]);
}

export function scaleRecipe(recipe, portions) {
    const scaledRecipe = {};
    Object.keys(recipe).forEach((ingredient) => {
        scaledRecipe[ingredient] = recipe[ingredient] * (portions / 2);
    });
    return scaledRecipe;
}