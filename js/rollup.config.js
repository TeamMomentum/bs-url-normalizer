import commonjs from '@rollup/plugin-commonjs';
import nodeResolve from '@rollup/plugin-node-resolve';
import json from '@rollup/plugin-json';
import { terser } from 'rollup-plugin-terser';

export default {
  output: {
    format: 'umd',
    name: 'urlnorm',
  },

  plugins: [
    nodeResolve({ preferBuiltins: false }),
    commonjs(),
    json(),
    terser(),
  ],
};
