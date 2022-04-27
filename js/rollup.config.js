import commonjs from 'rollup-plugin-commonjs';
import nodeResolve from 'rollup-plugin-node-resolve';

export default {
  output: {
    format: 'umd',
    name: 'urlnorm',
  },

  plugins: [nodeResolve(), commonjs()],
};
