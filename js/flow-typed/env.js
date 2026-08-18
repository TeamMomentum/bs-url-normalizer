// Local library definitions for Flow.
// Recent Flow versions no longer bundle Node/BOM libdefs and do not
// resolve untyped node_modules packages, so declare them here.

declare var window: any;
declare var process: any;

declare class URL {
  constructor(input: any, base?: string | URL): void;
  protocol: string;
  hostname: string;
  port: string;
  pathname: string;
  search: string;
  hash: string;
  href: string;
}

declare module 'assert' {
  declare module.exports: any;
}

declare module 'fs' {
  declare module.exports: any;
}

declare module 'path' {
  declare module.exports: any;
}

declare module 'mocha' {
  declare module.exports: any;
}

declare module 'url-parse' {
  declare module.exports: any;
}

declare module 'tr46' {
  declare module.exports: any;
}
