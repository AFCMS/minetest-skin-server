module.exports = {
  extends: ["stylelint-config-standard"],
  customSyntax: "postcss-scss",
  rules: {
    "at-rule-no-unknown": [
      true,
      {
        ignoreAtRules: [
          "tailwind",
          "apply",
          "variants",
          "responsive",
          "screen",
          "layer",
        ],
      },
    ],
    indentation: null,
    "value-list-comma-newline-after": null,
    "declaration-colon-newline-after": null,
  },
};
