module.exports = {
  extends: 'standard',
  env: {
    browser: true,
    node: true
  },
  rules: {
    complexity: 'error',
    'max-depth': 'error',
    'max-lines-per-function': [
      'error',
      { skipBlankLines: true, skipComments: true }
    ],
    'max-nested-callbacks': 'error',
    'max-statements': ['error', 25],

    semi: 'off',
    'space-before-function-paren': 'off',
    indent: 'off'
  }
};
