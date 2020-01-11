(function(o){function c(c){for(var r,t,s=c[0],n=c[1],i=c[2],f=0,d=[];f<s.length;f++)t=s[f],Object.prototype.hasOwnProperty.call(l,t)&&l[t]&&d.push(l[t][0]),l[t]=0;for(r in n)Object.prototype.hasOwnProperty.call(n,r)&&(o[r]=n[r]);u&&u(c);while(d.length)d.shift()();return a.push.apply(a,i||[]),e()}function e(){for(var o,c=0;c<a.length;c++){for(var e=a[c],r=!0,s=1;s<e.length;s++){var n=e[s];0!==l[n]&&(r=!1)}r&&(a.splice(c--,1),o=t(t.s=e[0]))}return o}var r={},l={app:0},a=[];function t(c){if(r[c])return r[c].exports;var e=r[c]={i:c,l:!1,exports:{}};return o[c].call(e.exports,e,e.exports,t),e.l=!0,e.exports}t.m=o,t.c=r,t.d=function(o,c,e){t.o(o,c)||Object.defineProperty(o,c,{enumerable:!0,get:e})},t.r=function(o){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(o,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(o,"__esModule",{value:!0})},t.t=function(o,c){if(1&c&&(o=t(o)),8&c)return o;if(4&c&&"object"===typeof o&&o&&o.__esModule)return o;var e=Object.create(null);if(t.r(e),Object.defineProperty(e,"default",{enumerable:!0,value:o}),2&c&&"string"!=typeof o)for(var r in o)t.d(e,r,function(c){return o[c]}.bind(null,r));return e},t.n=function(o){var c=o&&o.__esModule?function(){return o["default"]}:function(){return o};return t.d(c,"a",c),c},t.o=function(o,c){return Object.prototype.hasOwnProperty.call(o,c)},t.p="/";var s=window["webpackJsonp"]=window["webpackJsonp"]||[],n=s.push.bind(s);s.push=c,s=s.slice();for(var i=0;i<s.length;i++)c(s[i]);var u=n;a.push([0,"chunk-vendors"]),e()})({0:function(o,c,e){o.exports=e("56d7")},"034f":function(o,c,e){"use strict";var r=e("85ec"),l=e.n(r);l.a},"0898":function(o,c,e){"use strict";var r=e("8631"),l=e.n(r);l.a},3871:function(o,c,e){},"56d7":function(o,c,e){"use strict";e.r(c);e("e260"),e("e6cf"),e("cca6"),e("a79d");var r=e("2b0e"),l=function(){var o=this,c=o.$createElement,e=o._self._c||c;return e("div",{attrs:{id:"app"}},[e("SearchBar",{attrs:{selectedLanguages:o.selectedLanguages},on:{"search-issues":o.searchIssues,"select-language":o.selectLanguage}}),e("Issues",{attrs:{issues:o.issues}})],1)},a=[],t=(e("d3b7"),e("159b"),function(){var o=this,c=o.$createElement,e=o._self._c||c;return e("div",{staticClass:"search-container"},[e("div",{staticClass:"left-whitespace"}),e("div",{staticClass:"searchbar-inner-container"},[e("div",[e("div",{staticClass:"lang-selection"},o._l(o.selectedLanguages,(function(c){return e("div",{key:c.Language,staticClass:"lang-filter",class:{"selected-lang-filter":c.IsSelected},on:{click:function(e){return o.select(c)}}},[e("div",{staticClass:"lang-filter-text"},[o._v(o._s(c.Language))])])})),0)]),e("div",{staticClass:"search-button-parent"},[e("div",{staticClass:"search-button",on:{click:function(c){return o.searchIssues()}}},[o._v("Search")])])]),e("div",{staticClass:"right-whitespace"})])}),s=[],n={props:["selectedLanguages"],methods:{searchIssues:function(){this.$emit("search-issues")},select:function(o){this.$emit("select-language",o)}}},i=n,u=(e("a523"),e("2877")),f=Object(u["a"])(i,t,s,!1,null,"3314127d",null),d=f.exports,p=function(){var o=this,c=o.$createElement,e=o._self._c||c;return e("div",o._l(o.issues,(function(o){return e("div",{key:o.Title,staticClass:"issues-container"},[e("div",{staticClass:"left-whitespace"}),e("div",{staticClass:"issue-inner-container"},[e("IssueItem",{attrs:{issue:o}})],1),e("div",{staticClass:"right-whitespace"})])})),0)},g=[],b=function(){var o=this,c=o.$createElement,e=o._self._c||c;return e("div",{staticClass:"issue-item",on:{click:function(c){return o.openRepo()}}},[e("div",{staticClass:"issue-lang"},[e("span",{staticClass:"lang-circle",style:{background:o.issue.LangColor}}),e("span",{staticClass:"lang-text"},[o._v(o._s(o.issue.Language))])]),e("div",{staticClass:"issue-title"},[e("p",[o._v(o._s(o.issue.Title))])]),e("div",{staticClass:"issue-repo"},[e("span",{staticClass:"issue-repo-span"},[o._v(o._s(o.issue.Repo))])])])},h=[],C={name:"IssueItem",props:["issue"],methods:{openRepo:function(){var o=window.open("https://github.com/"+this.issue.Repo+"/issues/"+this.issue.Number);o.focus()}}},L=C,S=(e("a140"),Object(u["a"])(L,b,h,!1,null,"5d54d8d9",null)),v=S.exports,m={name:"Issues",components:{IssueItem:v},props:["issues"]},A=m,F=(e("0898"),Object(u["a"])(A,p,g,!1,null,"5c2b664d",null)),y=F.exports,B=e("bd22"),E=e("9f29"),P=e("17fb"),D=e.n(P),I={name:"app",components:{Issues:y,SearchBar:d},methods:{getBaseUrl:function(){return"https://cloud.fixthepla.net"},searchIssues:function(){var o=this,c="";this.selectedLanguages.forEach((function(o){o.IsSelected&&(c+=""==c?"?":"&",c+="language="+encodeURIComponent(o.Language))})),fetch(this.getBaseUrl()+"/issues"+c).then((function(c){c.json().then((function(c){o.issues=[],c.Issues.forEach((function(c){var e=D.a.some(o.issues,(function(o){return D.a.isEqual(o,c)}));e||o.issues.push(c)})),o.issues.forEach((function(c){c.LangColor=c.Language in o.languageColors?o.languageColors[c.Language].color:"#ffffff"}))})).catch((function(o){return console.log(o)}))})).catch((function(o){return console.log(o)}))},selectLanguage:function(o){this.selectedLanguages.forEach((function(c){c.Language==o.Language&&(c.IsSelected=!c.IsSelected)})),this.searchIssues()}},data:function(){return{issues:[],languageColors:B,mainLanguages:E,selectedLanguages:[]}},created:function(){var o=this;this.searchIssues(),this.mainLanguages.forEach((function(c){o.selectedLanguages.push({Language:c,IsSelected:!1})}))}},k=I,M=(e("034f"),Object(u["a"])(k,l,a,!1,null,null,null)),w=M.exports;r["a"].config.productionTip=!1,new r["a"]({render:function(o){return o(w)}}).$mount("#app")},"85ec":function(o,c,e){},8631:function(o,c,e){},"9b7f":function(o,c,e){},"9f29":function(o){o.exports=JSON.parse('["JavaScript","Python","Java","PHP","C#","C++","TypeScript","Shell","C","Ruby","Go","Swift","Rust","HTML","Dart"]')},a140:function(o,c,e){"use strict";var r=e("9b7f"),l=e.n(r);l.a},a523:function(o,c,e){"use strict";var r=e("3871"),l=e.n(r);l.a},bd22:function(o){o.exports=JSON.parse('{"1C Enterprise":{"color":"#814CCC"},"ABAP":{"color":"#E8274B"},"ActionScript":{"color":"#882B0F"},"Ada":{"color":"#02f88c"},"Agda":{"color":"#315665"},"AGS Script":{"color":"#B9D9FF"},"Alloy":{"color":"#64C800"},"AMPL":{"color":"#E6EFBB"},"AngelScript":{"color":"#C7D7DC"},"ANTLR":{"color":"#9DC3FF"},"API Blueprint":{"color":"#2ACCA8"},"APL":{"color":"#5A8164"},"AppleScript":{"color":"#101F1F"},"Arc":{"color":"#aa2afe"},"ASP":{"color":"#6a40fd"},"AspectJ":{"color":"#a957b0"},"Assembly":{"color":"#6E4C13"},"Asymptote":{"color":"#4a0c0c"},"ATS":{"color":"#1ac620"},"AutoHotkey":{"color":"#6594b9"},"AutoIt":{"color":"#1C3552"},"Ballerina":{"color":"#FF5000"},"Batchfile":{"color":"#C1F12E"},"BlitzMax":{"color":"#cd6400"},"Boo":{"color":"#d4bec1"},"Brainfuck":{"color":"#2F2530"},"C":{"color":"#555555"},"C#":{"color":"#178600"},"C++":{"color":"#f34b7d"},"Ceylon":{"color":"#dfa535"},"Chapel":{"color":"#8dc63f"},"Cirru":{"color":"#ccccff"},"Clarion":{"color":"#db901e"},"Clean":{"color":"#3F85AF"},"Click":{"color":"#E4E6F3"},"Clojure":{"color":"#db5855"},"CoffeeScript":{"color":"#244776"},"ColdFusion":{"color":"#ed2cd6"},"Common Lisp":{"color":"#3fb68b"},"Common Workflow Language":{"color":"#B5314C"},"Component Pascal":{"color":"#B0CE4E"},"Crystal":{"color":"#000100"},"CSS":{"color":"#563d7c"},"Cuda":{"color":"#3A4E3A"},"D":{"color":"#ba595e"},"Dart":{"color":"#00B4AB"},"DataWeave":{"color":"#003a52"},"DM":{"color":"#447265"},"Dockerfile":{"color":"#384d54"},"Dogescript":{"color":"#cca760"},"Dylan":{"color":"#6c616e"},"E":{"color":"#ccce35"},"eC":{"color":"#913960"},"ECL":{"color":"#8a1267"},"Eiffel":{"color":"#946d57"},"Elixir":{"color":"#6e4a7e"},"Elm":{"color":"#60B5CC"},"Emacs Lisp":{"color":"#c065db"},"EmberScript":{"color":"#FFF4F3"},"EQ":{"color":"#a78649"},"Erlang":{"color":"#B83998"},"F#":{"color":"#b845fc"},"F*":{"color":"#572e30"},"Factor":{"color":"#636746"},"Fancy":{"color":"#7b9db4"},"Fantom":{"color":"#14253c"},"FLUX":{"color":"#88ccff"},"Forth":{"color":"#341708"},"Fortran":{"color":"#4d41b1"},"FreeMarker":{"color":"#0050b2"},"Frege":{"color":"#00cafe"},"Game Maker Language":{"color":"#71b417"},"GDScript":{"color":"#355570"},"Genie":{"color":"#fb855d"},"Gherkin":{"color":"#5B2063"},"Glyph":{"color":"#c1ac7f"},"Gnuplot":{"color":"#f0a9f0"},"Go":{"color":"#00ADD8"},"Golo":{"color":"#88562A"},"Gosu":{"color":"#82937f"},"Grammatical Framework":{"color":"#79aa7a"},"Groovy":{"color":"#e69f56"},"Hack":{"color":"#878787"},"Harbour":{"color":"#0e60e3"},"Haskell":{"color":"#5e5086"},"Haxe":{"color":"#df7900"},"HiveQL":{"color":"#dce200"},"HTML":{"color":"#e34c26"},"Hy":{"color":"#7790B2"},"IDL":{"color":"#a3522f"},"Idris":{"color":"#b30000"},"Io":{"color":"#a9188d"},"Ioke":{"color":"#078193"},"Isabelle":{"color":"#FEFE00"},"J":{"color":"#9EEDFF"},"Java":{"color":"#b07219"},"JavaScript":{"color":"#f1e05a"},"Jolie":{"color":"#843179"},"JSONiq":{"color":"#40d47e"},"Jsonnet":{"color":"#0064bd"},"Julia":{"color":"#a270ba"},"Jupyter Notebook":{"color":"#DA5B0B"},"Kotlin":{"color":"#F18E33"},"KRL":{"color":"#28430A"},"Lasso":{"color":"#999999"},"Lex":{"color":"#DBCA00"},"LFE":{"color":"#4C3023"},"LiveScript":{"color":"#499886"},"LLVM":{"color":"#185619"},"LOLCODE":{"color":"#cc9900"},"LookML":{"color":"#652B81"},"LSL":{"color":"#3d9970"},"Lua":{"color":"#000080"},"Makefile":{"color":"#427819"},"Mask":{"color":"#f97732"},"MATLAB":{"color":"#e16737"},"Max":{"color":"#c4a79c"},"MAXScript":{"color":"#00a6a6"},"mcfunction":{"color":"#E22837"},"Mercury":{"color":"#ff2b2b"},"Meson":{"color":"#007800"},"Metal":{"color":"#8f14e9"},"Mirah":{"color":"#c7a938"},"Modula-3":{"color":"#223388"},"MQL4":{"color":"#62A8D6"},"MQL5":{"color":"#4A76B8"},"MTML":{"color":"#b7e1f4"},"NCL":{"color":"#28431f"},"Nearley":{"color":"#990000"},"Nemerle":{"color":"#3d3c6e"},"nesC":{"color":"#94B0C7"},"NetLinx":{"color":"#0aa0ff"},"NetLinx+ERB":{"color":"#747faa"},"NetLogo":{"color":"#ff6375"},"NewLisp":{"color":"#87AED7"},"Nextflow":{"color":"#3ac486"},"Nim":{"color":"#37775b"},"Nit":{"color":"#009917"},"Nix":{"color":"#7e7eff"},"Nu":{"color":"#c9df40"},"Objective-C":{"color":"#438eff"},"Objective-C++":{"color":"#6866fb"},"Objective-J":{"color":"#ff0c5a"},"OCaml":{"color":"#3be133"},"Omgrofl":{"color":"#cabbff"},"ooc":{"color":"#b0b77e"},"Opal":{"color":"#f7ede0"},"Oxygene":{"color":"#cdd0e3"},"Oz":{"color":"#fab738"},"P4":{"color":"#7055b5"},"Pan":{"color":"#cc0000"},"Papyrus":{"color":"#6600cc"},"Parrot":{"color":"#f3ca0a"},"Pascal":{"color":"#E3F171"},"Pawn":{"color":"#dbb284"},"Pep8":{"color":"#C76F5B"},"Perl":{"color":"#0298c3"},"Perl 6":{"color":"#0000fb"},"PHP":{"color":"#4F5D95"},"PigLatin":{"color":"#fcd7de"},"Pike":{"color":"#005390"},"PLSQL":{"color":"#dad8d8"},"PogoScript":{"color":"#d80074"},"PostScript":{"color":"#da291c"},"PowerBuilder":{"color":"#8f0f8d"},"PowerShell":{"color":"#012456"},"Processing":{"color":"#0096D8"},"Prolog":{"color":"#74283c"},"Propeller Spin":{"color":"#7fa2a7"},"Puppet":{"color":"#302B6D"},"PureBasic":{"color":"#5a6986"},"PureScript":{"color":"#1D222D"},"Python":{"color":"#3572A5"},"q":{"color":"#0040cd"},"QML":{"color":"#44a51c"},"Quake":{"color":"#882233"},"R":{"color":"#198CE7"},"Racket":{"color":"#3c5caa"},"Ragel":{"color":"#9d5200"},"RAML":{"color":"#77d9fb"},"Rascal":{"color":"#fffaa0"},"Rebol":{"color":"#358a5b"},"Red":{"color":"#f50000"},"Ren\'Py":{"color":"#ff7f7f"},"Ring":{"color":"#2D54CB"},"Roff":{"color":"#ecdebe"},"Rouge":{"color":"#cc0088"},"Ruby":{"color":"#701516"},"RUNOFF":{"color":"#665a4e"},"Rust":{"color":"#dea584"},"SaltStack":{"color":"#646464"},"SAS":{"color":"#B34936"},"Scala":{"color":"#c22d40"},"Scheme":{"color":"#1e4aec"},"sed":{"color":"#64b970"},"Self":{"color":"#0579aa"},"Shell":{"color":"#89e051"},"Shen":{"color":"#120F14"},"Slash":{"color":"#007eff"},"Slice":{"color":"#003fa2"},"Smalltalk":{"color":"#596706"},"Solidity":{"color":"#AA6746"},"SourcePawn":{"color":"#5c7611"},"SQF":{"color":"#3F3F3F"},"Squirrel":{"color":"#800000"},"SRecode Template":{"color":"#348a34"},"Stan":{"color":"#b2011d"},"Standard ML":{"color":"#dc566d"},"SuperCollider":{"color":"#46390b"},"Swift":{"color":"#ffac45"},"SystemVerilog":{"color":"#DAE1C2"},"Tcl":{"color":"#e4cc98"},"Terra":{"color":"#00004c"},"TeX":{"color":"#3D6117"},"TI Program":{"color":"#A0AA87"},"Turing":{"color":"#cf142b"},"TypeScript":{"color":"#2b7489"},"UnrealScript":{"color":"#a54c4d"},"Vala":{"color":"#fbe5cd"},"VCL":{"color":"#148AA8"},"Verilog":{"color":"#b2b7f8"},"VHDL":{"color":"#adb2cb"},"Vim script":{"color":"#199f4b"},"Visual Basic":{"color":"#945db7"},"Volt":{"color":"#1F1F1F"},"Vue":{"color":"#2c3e50"},"wdl":{"color":"#42f1f4"},"WebAssembly":{"color":"#04133b"},"wisp":{"color":"#7582D1"},"X10":{"color":"#4B6BEF"},"xBase":{"color":"#403a40"},"XC":{"color":"#99DA07"},"XQuery":{"color":"#5232e7"},"XSLT":{"color":"#EB8CEB"},"Yacc":{"color":"#4B6C4B"},"YARA":{"color":"#220000"},"YASnippet":{"color":"#32AB90"},"ZAP":{"color":"#0d665e"},"Zephir":{"color":"#118f9e"},"Zig":{"color":"#ec915c"},"ZIL":{"color":"#dc75e5"}}')}});
//# sourceMappingURL=app.4744e794.js.map