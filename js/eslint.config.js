import neostandard from 'neostandard';

export default [
  ...neostandard({
    // Formatting is handled by Prettier; disable style rules
    // (successor of the old comma-dangle/indent/semi/... "off" settings).
    noStyle: true,
    env: ['browser', 'node'],
  }),
  {
    rules: {
      complexity: 'error',
      'max-depth': 'error',
      'max-lines-per-function': [
        'error',
        { skipBlankLines: true, skipComments: true },
      ],
      'max-nested-callbacks': 'error',
      'max-statements': ['error', 25],

      'no-var': 'off',
      'object-shorthand': 'off',
    },
  },
  {
    ignores: ['build/', 'flow-typed/', 'node_modules/'],
  },
];
