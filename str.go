// Credits:
// https://github.com/jprichardson/string.js
// https://github.com/anbinh/dna/blob/master/string.go

// Package str is a set of useful string helpers which can be composed
// into a pipeline.
package str

import (
	//	"log"
	"regexp"
	"strings"
)

var spacesRe = regexp.MustCompile("[\\s\\xA0]+")
var beginEndSpacesRe = regexp.MustCompile("^\\s+|\\s+$")
var spaceUnderscoreRe = regexp.MustCompile("[_\\s]+")
var capitalsRe = regexp.MustCompile("([A-Z])")
var dashesRe = regexp.MustCompile("-+")
var notDigitsRe = regexp.MustCompile(`[^0-9]`)
var isAlphaRe = regexp.MustCompile(`[^a-z\xC0-\xFF]`)
var isAlphaNumericRe = regexp.MustCompile(`[^0-9a-z\xC0-\xFF]`)

// Between extracts a string between left and right strings.
func Between(s, left, right string) string {
	l := len(left)
	startPos := strings.Index(s, left)
	endPos := IndexOf(s, right, startPos+l)
	//log.Printf("%s: left %s right %s start %d end %d", s, left, right, startPos+l, endPos)
	if endPos < 0 {
		return ""
	} else if right == "" {
		return s[endPos:]
	} else {
		return s[startPos+l : endPos]
	}
}

// BetweenF is the filter form for Between.
func BetweenF(left, right string) func(string) string {
	return func(s string) string {
		return Between(s, left, right)
	}
}

// Camelize return new string which removes any underscores or dashes and convert a string into camel casing.
func Camelize(s string) string {
	r := regexp.MustCompile(`(\-|_|\s)+(.)?`)
	return r.ReplaceAllStringFunc(s, func(val string) string {
		val = strings.ToUpper(val)
		val = ReplaceWithRegexp(val, `(\-|_|\s)+`, "")
		return val
	})
}

//     capitalize: function() {
//       return new this.constructor(this.s.substr(0, 1).toUpperCase() + this.s.substring(1).toLowerCase());
//     },
//

// CharAt returns a string from the character at the specified position.
func CharAt(s string, index int) string {
	l := len(s)
	shortcut := index < 0 || index > l-1 || l == 0
	if shortcut {
		return ""
	}
	return s[index : index+1]
}

// CharAtF is the filter form of CharAt.
func CharAtF(index int) func(string) string {
	return func(s string) string {
		return CharAt(s, index)
	}
}

// ChompLeft removes prefix at the start of a string.
func ChompLeft(s, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s
}

// ChompLeft is the filter form of ChompLeft.
func ChompLeftF(prefix string) func(string) string {
	return func(s string) string {
		return ChompLeft(s, prefix)
	}
}

// ChompRight removes suffix from end of s.
func ChompRight(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return s[:len(s)-len(suffix)]
	}
	return s
}

// ChompRightF is the filter form of ChompRight.
func ChompRightF(suffix string) func(string) string {
	return func(s string) string {
		return ChompRight(s, suffix)
	}
}

// Clean compresses all adjacent whitespace to a single space and trims s.
func Clean(s string) string {
	s = spacesRe.ReplaceAllString(s, " ")
	s = beginEndSpacesRe.ReplaceAllString(s, "")
	return s
}

//     count: function(ss) {
//       var count = 0
//         , pos = this.s.indexOf(ss)
//
//       while (pos >= 0) {
//         count += 1
//         pos = this.s.indexOf(ss, pos + 1)
//       }
//
//       return count
//     },
//

// Dasherize  converts a camel cased string into a string delimited by dashes.
func Dasherize(s string) string {
	prefixed := false
	if strings.HasPrefix(s, "-") {
		prefixed = true
	}

	s = strings.TrimSpace(s)
	s = spaceUnderscoreRe.ReplaceAllString(s, "-")
	s = capitalsRe.ReplaceAllString(s, "-$1")
	s = dashesRe.ReplaceAllString(s, "-")
	s = strings.ToLower(s)
	if !prefixed && strings.HasPrefix(s, "-") {
		s = s[1:]
	}
	return s
}

//     decodeHtmlEntities: function() { //https://github.com/substack/node-ent/blob/master/index.js
//       var s = this.s;
//       s = s.replace(/&#(\d+);?/g, function (_, code) {
//         return String.fromCharCode(code);
//       })
//       .replace(/&#[xX]([A-Fa-f0-9]+);?/g, function (_, hex) {
//         return String.fromCharCode(parseInt(hex, 16));
//       })
//       .replace(/&([^;\W]+;?)/g, function (m, e) {
//         var ee = e.replace(/;$/, '');
//         var target = ENTITIES[e] || (e.match(/;$/) && ENTITIES[ee]);
//
//         if (typeof target === 'number') {
//           return String.fromCharCode(target);
//         }
//         else if (typeof target === 'string') {
//           return target;
//         }
//         else {
//           return m;
//         }
//       })
//
//       return new this.constructor(s);
//     },
//
//     escapeHTML: function() { //from underscore.string
//       return new this.constructor(this.s.replace(/[&<>"']/g, function(m){ return '&' + reversedEscapeChars[m] + ';'; }));
//     },
//

// EnsurePrefix ensures s starts with prefix.
func EnsurePrefix(s, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		return s
	}
	return prefix + s
}

// EnsurePrefixF is the filter form of EnsurePrefix.
func EnsurePrefixF(prefix string) func(string) string {
	return func(s string) string {
		return EnsurePrefix(s, prefix)
	}
}

// EnsureSuffix ensures s ends with suffix.
func EnsureSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return s
	}
	return s + suffix
}

// EnsureSuffixF is the filter version of EnsureSuffix.
func EnsureSuffixF(suffix string) func(string) string {
	return func(s string) string {
		return EnsureSuffix(s, suffix)
	}
}

//     humanize: function() {
//       if (this.s === null || this.s === undefined)
//         return new this.constructor('')
//       var s = this.underscore().replace(/_id$/,'').replace(/_/g, ' ').trim().capitalize()
//       return new this.constructor(s)
//     }

// IndexOf finds the index of needle in s starting from start.
func IndexOf(s string, needle string, start int) int {
	l := len(s)
	if needle == "" {
		if start < 0 {
			return 0
		} else if start < l {
			return start
		} else {
			return l
		}
	}
	if start < 0 || start > l-1 {
		return -1
	}
	pos := strings.Index(s[start:], needle)
	if pos == -1 {
		return -1
	} else {
		return start + pos
	}
}

// IsAlpha returns true if a string contains only letters from ASCII (a-z,A-Z). Other letters from other languages are not supported.
func IsAlpha(s string) bool {
	return !isAlphaRe.MatchString(strings.ToLower(s))
}

// IsAlphaNumeric returns true if a string contains letters and digits.
func IsAlphaNumeric(s string) bool {
	return !isAlphaNumericRe.MatchString(strings.ToLower(s))
}

// IsLower returns true if s comprised of all lower case characters.
func IsLower(s string) bool {
	return IsAlpha(s) && s == strings.ToLower(s)
}

// IsNumeric returns true if a string contains only digits from 0-9. Other digits not in Latin (such as Arabic) are not currently supported.
func IsNumeric(s string) bool {
	return !notDigitsRe.MatchString(s)
}

// IsUpper returns true if s contains all upper case chracters.
func IsUpper(s string) bool {
	return IsAlpha(s) && s == strings.ToUpper(s)
}

//
//     isAlpha: function() {
//       return !/[^a-z\xC0-\xFF]/.test(this.s.toLowerCase());
//     },
//
//     isAlphaNumeric: function() {
//       return !/[^0-9a-z\xC0-\xFF]/.test(this.s.toLowerCase());
//     },
//
//     isEmpty: function() {
//       return this.s === null || this.s === undefined ? true : /^[\s\xa0]*$/.test(this.s);
//     },
//
//     isLower: function() {
//       return this.isAlpha() && this.s.toLowerCase() === this.s;
//     },
//
//     isNumeric: function() {
//       return !/[^0-9]/.test(this.s);
//     },
//
//     isUpper: function() {
//       return this.isAlpha() && this.s.toUpperCase() === this.s;
//     },
//
//     left: function(N) {
//       if (N >= 0) {
//         var s = this.s.substr(0, N);
//         return new this.constructor(s);
//       } else {
//         return this.right(-N);
//       }
//     },
//
//     lines: function() { //convert windows newlines to unix newlines then convert to an Array of lines
//       return this.replaceAll('\r\n', '\n').s.split('\n');
//     },
//
//     pad: function(len, ch) { //https://github.com/component/pad
//       if (ch == null) ch = ' ';
//       if (this.s.length >= len) return new this.constructor(this.s);
//       len = len - this.s.length;
//       var left = Array(Math.ceil(len / 2) + 1).join(ch);
//       var right = Array(Math.floor(len / 2) + 1).join(ch);
//       return new this.constructor(left + this.s + right);
//     },
//
//     padLeft: function(len, ch) { //https://github.com/component/pad
//       if (ch == null) ch = ' ';
//       if (this.s.length >= len) return new this.constructor(this.s);
//       return new this.constructor(Array(len - this.s.length + 1).join(ch) + this.s);
//     },
//
//     padRight: function(len, ch) {
//       if (ch == null) ch = ' ';
//       if (this.s.length >= len) return new this.constructor(this.s);
//       return new this.constructor(this.s + Array(len - this.s.length + 1).join(ch));
//     },
//
//     parseCSV: function(delimiter, qualifier, escape, lineDelimiter) { //try to parse no matter what
//       delimiter = delimiter || ',';
//       escape = escape || '\\'
//       if (typeof qualifier == 'undefined')
//         qualifier = '"';
//
//       var i = 0, fieldBuffer = [], fields = [], len = this.s.length, inField = false, self = this;
//       var ca = function(i){return self.s.charAt(i)};
//       if (typeof lineDelimiter !== 'undefined') var rows = [];
//
//       if (!qualifier)
//         inField = true;
//
//       while (i < len) {
//         var current = ca(i);
//         switch (current) {
//           case escape:
//           //fix for issues #32 and #35
//           if (inField && ((escape !== qualifier) || ca(i+1) === qualifier)) {
//               i += 1;
//               fieldBuffer.push(ca(i));
//               break;
//           }
//           if (escape !== qualifier) break;
//           case qualifier:
//             inField = !inField;
//             break;
//           case delimiter:
//             if (inField && qualifier)
//               fieldBuffer.push(current);
//             else {
//               fields.push(fieldBuffer.join(''))
//               fieldBuffer.length = 0;
//             }
//             break;
//           case lineDelimiter:
//             if (inField) {
//                 fieldBuffer.push(current);
//             } else {
//                 if (rows) {
//                     fields.push(fieldBuffer.join(''))
//                     rows.push(fields);
//                     fields = [];
//                     fieldBuffer.length = 0;
//                 }
//             }
//             break;
//           default:
//             if (inField)
//               fieldBuffer.push(current);
//             break;
//         }
//         i += 1;
//       }
//
//       fields.push(fieldBuffer.join(''));
//       if (rows) {
//         rows.push(fields);
//         return rows;
//       }
//       return fields;
//     },

// Match returns true if patterns matches the string
func Match(s, pattern string) bool {
	r := regexp.MustCompile(pattern)
	return r.MatchString(s)
}

// Pipe pipes s through one or more string filters.
func Pipe(s string, funcs ...func(string) string) string {
	for _, fn := range funcs {
		s = fn(s)
	}
	return s
}

//
//     replaceAll: function(ss, r) {
//       //var s = this.s.replace(new RegExp(ss, 'g'), r);
//       var s = this.s.split(ss).join(r)
//       return new this.constructor(s);
//     },

// ReplaceWithRegexp replaces string with regexp string.
// ReplaceWithRegexp returns a copy of src, replacing matches of the Regexp with the replacement string repl. Inside repl, $ signs are interpreted as in Expand, so for instance $1 represents the text of the first submatch.
func ReplaceWithRegexp(s, pattern, repl string) string {
	r := regexp.MustCompile(pattern)
	return r.ReplaceAllString(s, repl)
}

//
//     right: function(N) {
//       if (N >= 0) {
//         var s = this.s.substr(this.s.length - N, N);
//         return new this.constructor(s);
//       } else {
//         return this.left(-N);
//       }
//     },
//
//     setValue: function (s) {
// 	  initialize(this, s);
// 	  return this;
//     },

// Slice slices a string. If end is negative then it is the from the end
// of the string.
func Slice(s string, start, end int) string {
	if end > -1 {
		return s[start:end]
	} else {
		L := len(s)
		if L+end > 0 {
			return s[start : L-end]
		}
		return s[start:]
	}
}

// SliceF is the filter for Slice.
func SliceF(start, end int) func(string) string {
	return func(s string) string {
		return Slice(s, start, end)
	}
}

//
//     slugify: function() {
//       var sl = (new S(this.s.replace(/[^\w\s-]/g, '').toLowerCase())).dasherize().s;
//       if (sl.charAt(0) === '-')
//         sl = sl.substr(1);
//       return new this.constructor(sl);
//     },
//
//     startsWith: function(prefix) {
//       return this.s.lastIndexOf(prefix, 0) === 0;
//     },
//
//     stripPunctuation: function() {
//       //return new this.constructor(this.s.replace(/[\.,-\/#!$%\^&\*;:{}=\-_`~()]/g,""));
//       return new this.constructor(this.s.replace(/[^\w\s]|_/g, "").replace(/\s+/g, " "));
//     },
//
//     stripTags: function() { //from sugar.js
//       var s = this.s, args = arguments.length > 0 ? arguments : [''];
//       multiArgs(args, function(tag) {
//         s = s.replace(RegExp('<\/?' + tag + '[^<>]*>', 'gi'), '');
//       });
//       return new this.constructor(s);
//     },
//
//     template: function(values, opening, closing) {
//       var s = this.s
//       var opening = opening || Export.TMPL_OPEN
//       var closing = closing || Export.TMPL_CLOSE
//
//       var open = opening.replace(/[-[\]()*\s]/g, "\\$&").replace(/\$/g, '\\$')
//       var close = closing.replace(/[-[\]()*\s]/g, "\\$&").replace(/\$/g, '\\$')
//       var r = new RegExp(open + '(.+?)' + close, 'g')
//         //, r = /\{\{(.+?)\}\}/g
//       var matches = s.match(r) || [];
//
//       matches.forEach(function(match) {
//         var key = match.substring(opening.length, match.length - closing.length);//chop {{ and }}
//         if (typeof values[key] != 'undefined')
//           s = s.replace(match, values[key]);
//       });
//       return new this.constructor(s);
//     },
//
//     times: function(n) {
//       return new this.constructor(new Array(n + 1).join(this.s));
//     },
//
//     toBoolean: function() {
//       if (typeof this.orig === 'string') {
//         var s = this.s.toLowerCase();
//         return s === 'true' || s === 'yes' || s === 'on' || s === '1';
//       } else
//         return this.orig === true || this.orig === 1;
//     },
//
//     toFloat: function(precision) {
//       var num = parseFloat(this.s)
//       if (precision)
//         return parseFloat(num.toFixed(precision))
//       else
//         return num
//     },
//
//     toInt: function() { //thanks Google
//       // If the string starts with '0x' or '-0x', parse as hex.
//       return /^\s*-?0x/i.test(this.s) ? parseInt(this.s, 16) : parseInt(this.s, 10)
//     },
//
//     trim: function() {
//       var s;
//       if (typeof __nsp.trim === 'undefined')
//         s = this.s.replace(/(^\s*|\s*$)/g, '')
//       else
//         s = this.s.trim()
//       return new this.constructor(s);
//     },
//
//     trimLeft: function() {
//       var s;
//       if (__nsp.trimLeft)
//         s = this.s.trimLeft();
//       else
//         s = this.s.replace(/(^\s*)/g, '');
//       return new this.constructor(s);
//     },
//
//     trimRight: function() {
//       var s;
//       if (__nsp.trimRight)
//         s = this.s.trimRight();
//       else
//         s = this.s.replace(/\s+$/, '');
//       return new this.constructor(s);
//     },
//
//     truncate: function(length, pruneStr) { //from underscore.string, author: github.com/rwz
//       var str = this.s;
//
//       length = ~~length;
//       pruneStr = pruneStr || '...';
//
//       if (str.length <= length) return new this.constructor(str);
//
//       var tmpl = function(c){ return c.toUpperCase() !== c.toLowerCase() ? 'A' : ' '; },
//         template = str.slice(0, length+1).replace(/.(?=\W*\w*$)/g, tmpl); // 'Hello, world' -> 'HellAA AAAAA'
//
//       if (template.slice(template.length-2).match(/\w\w/))
//         template = template.replace(/\s*\S+$/, '');
//       else
//         template = new S(template.slice(0, template.length-1)).trimRight().s;
//
//       return (template+pruneStr).length > str.length ? new S(str) : new S(str.slice(0, template.length)+pruneStr);
//     },
//
//     toCSV: function() {
//       var delim = ',', qualifier = '"', escape = '\\', encloseNumbers = true, keys = false;
//       var dataArray = [];
//
//       function hasVal(it) {
//         return it !== null && it !== '';
//       }
//
//       if (typeof arguments[0] === 'object') {
//         delim = arguments[0].delimiter || delim;
//         delim = arguments[0].separator || delim;
//         qualifier = arguments[0].qualifier || qualifier;
//         encloseNumbers = !!arguments[0].encloseNumbers;
//         escape = arguments[0].escape || escape;
//         keys = !!arguments[0].keys;
//       } else if (typeof arguments[0] === 'string') {
//         delim = arguments[0];
//       }
//
//       if (typeof arguments[1] === 'string')
//         qualifier = arguments[1];
//
//       if (arguments[1] === null)
//         qualifier = null;
//
//        if (this.orig instanceof Array)
//         dataArray  = this.orig;
//       else { //object
//         for (var key in this.orig)
//           if (this.orig.hasOwnProperty(key))
//             if (keys)
//               dataArray.push(key);
//             else
//               dataArray.push(this.orig[key]);
//       }
//
//       var rep = escape + qualifier;
//       var buildString = [];
//       for (var i = 0; i < dataArray.length; ++i) {
//         var shouldQualify = hasVal(qualifier)
//         if (typeof dataArray[i] == 'number')
//           shouldQualify &= encloseNumbers;
//
//         if (shouldQualify)
//           buildString.push(qualifier);
//
//         if (dataArray[i] !== null && dataArray[i] !== undefined) {
//           var d = new S(dataArray[i]).replaceAll(qualifier, rep).s;
//           buildString.push(d);
//         } else
//           buildString.push('')
//
//         if (shouldQualify)
//           buildString.push(qualifier);
//
//         if (delim)
//           buildString.push(delim);
//       }
//
//       //chop last delim
//       //console.log(buildString.length)
//       buildString.length = buildString.length - 1;
//       return new this.constructor(buildString.join(''));
//     },
//
//     toString: function() {
//       return this.s;
//     },
//
//     //#modified from https://github.com/epeli/underscore.string
//     underscore: function() {
//       var s = this.trim().s.replace(/([a-z\d])([A-Z]+)/g, '$1_$2').replace(/[-\s]+/g, '_').toLowerCase();
//       if ((new S(this.s.charAt(0))).isUpper()) {
//         s = '_' + s;
//       }
//       return new this.constructor(s);
//     },
//
//     unescapeHTML: function() { //from underscore.string
//       return new this.constructor(this.s.replace(/\&([^;]+);/g, function(entity, entityCode){
//         var match;
//
//         if (entityCode in escapeChars) {
//           return escapeChars[entityCode];
//         } else if (match = entityCode.match(/^#x([\da-fA-F]+)$/)) {
//           return String.fromCharCode(parseInt(match[1], 16));
//         } else if (match = entityCode.match(/^#(\d+)$/)) {
//           return String.fromCharCode(~~match[1]);
//         } else {
//           return entity;
//         }
//       }));
//     },
//
//     valueOf: function() {
//       return this.s.valueOf();
//     },
//
//     //#Added a New Function called wrapHTML.
//     wrapHTML: function (tagName, tagAttrs) {
//       var s = this.s, el = (tagName == null) ? 'span' : tagName, elAttr = '', wrapped = '';
//       if(typeof tagAttrs == 'object') for(var prop in tagAttrs) elAttr += ' ' + prop + '="' + tagAttrs[prop] + '"';
//       s = wrapped.concat('<', el, elAttr, '>', this, '</', el, '>');
//       return new this.constructor(s);
//     }
//   }

//////////////////////////////////////////////////////////////

//
//   var methodsAdded = [];
//   function extendPrototype() {
//     for (var name in __sp) {
//       (function(name){
//         var func = __sp[name];
//         if (!__nsp.hasOwnProperty(name)) {
//           methodsAdded.push(name);
//           __nsp[name] = function() {
//             String.prototype.s = this;
//             return func.apply(this, arguments);
//           }
//         }
//       })(name);
//     }
//   }
//
//   function restorePrototype() {
//     for (var i = 0; i < methodsAdded.length; ++i)
//       delete String.prototype[methodsAdded[i]];
//     methodsAdded.length = 0;
//   }
//
//
// /*************************************
// /* Attach Native JavaScript String Properties
// /*************************************/
//
//   var nativeProperties = getNativeStringProperties();
//   for (var name in nativeProperties) {
//     (function(name) {
//       var stringProp = __nsp[name];
//       if (typeof stringProp == 'function') {
//         //console.log(stringProp)
//         if (!__sp[name]) {
//           if (nativeProperties[name] === 'string') {
//             __sp[name] = function() {
//               //console.log(name)
//               return new this.constructor(stringProp.apply(this, arguments));
//             }
//           } else {
//             __sp[name] = stringProp;
//           }
//         }
//       }
//     })(name);
//   }
//
//
// /*************************************
// /* Function Aliases
// /*************************************/
//
//   __sp.repeat = __sp.times;
//   __sp.include = __sp.contains;
//   __sp.toInteger = __sp.toInt;
//   __sp.toBool = __sp.toBoolean;
//   __sp.decodeHTMLEntities = __sp.decodeHtmlEntities //ensure consistent casing scheme of 'HTML'
//
//
// //******************************************************************************
// // Set the constructor.  Without this, string.js objects are instances of
// // Object instead of S.
// //******************************************************************************
//
//   __sp.constructor = S;
//
//
// /*************************************
// /* Private Functions
// /*************************************/
//
//   function getNativeStringProperties() {
//     var names = getNativeStringPropertyNames();
//     var retObj = {};
//
//     for (var i = 0; i < names.length; ++i) {
//       var name = names[i];
//       var func = __nsp[name];
//       try {
//         var type = typeof func.apply('teststring', []);
//         retObj[name] = type;
//       } catch (e) {}
//     }
//     return retObj;
//   }
//
//   function getNativeStringPropertyNames() {
//     var results = [];
//     if (Object.getOwnPropertyNames) {
//       results = Object.getOwnPropertyNames(__nsp);
//       results.splice(results.indexOf('valueOf'), 1);
//       results.splice(results.indexOf('toString'), 1);
//       return results;
//     } else { //meant for legacy cruft, this could probably be made more efficient
//       var stringNames = {};
//       var objectNames = [];
//       for (var name in String.prototype)
//         stringNames[name] = name;
//
//       for (var name in Object.prototype)
//         delete stringNames[name];
//
//       //stringNames['toString'] = 'toString'; //this was deleted with the rest of the object names
//       for (var name in stringNames) {
//         results.push(name);
//       }
//       return results;
//     }
//   }
//
//   function Export(str) {
//     return new S(str);
//   };
//
//   //attach exports to StringJSWrapper
//   Export.extendPrototype = extendPrototype;
//   Export.restorePrototype = restorePrototype;
//   Export.VERSION = VERSION;
//   Export.TMPL_OPEN = '{{';
//   Export.TMPL_CLOSE = '}}';
//   Export.ENTITIES = ENTITIES;
//
//
//
// /*************************************
// /* Exports
// /*************************************/
//
//   if (typeof module !== 'undefined'  && typeof module.exports !== 'undefined') {
//     module.exports = Export;
//
//   } else {
//
//     if(typeof define === "function" && define.amd) {
//       define([], function() {
//         return Export;
//       });
//     } else {
//       window.S = Export;
//     }
//   }
//
//
// /*************************************
// /* 3rd Party Private Functions
// /*************************************/
//
//   //from sugar.js
//   function multiArgs(args, fn) {
//     var result = [], i;
//     for(i = 0; i < args.length; i++) {
//       result.push(args[i]);
//       if(fn) fn.call(args, args[i], i);
//     }
//     return result;
//   }
//
//   //from underscore.string
//   var escapeChars = {
//     lt: '<',
//     gt: '>',
//     quot: '"',
//     apos: "'",
//     amp: '&'
//   };
//
//   //from underscore.string
//   var reversedEscapeChars = {};
//   for(var key in escapeChars){ reversedEscapeChars[escapeChars[key]] = key; }
//
//   ENTITIES = {
//     "amp" : "&",
//     "gt" : ">",
//     "lt" : "<",
//     "quot" : "\"",
//     "apos" : "'",
//     "AElig" : 198,
//     "Aacute" : 193,
//     "Acirc" : 194,
//     "Agrave" : 192,
//     "Aring" : 197,
//     "Atilde" : 195,
//     "Auml" : 196,
//     "Ccedil" : 199,
//     "ETH" : 208,
//     "Eacute" : 201,
//     "Ecirc" : 202,
//     "Egrave" : 200,
//     "Euml" : 203,
//     "Iacute" : 205,
//     "Icirc" : 206,
//     "Igrave" : 204,
//     "Iuml" : 207,
//     "Ntilde" : 209,
//     "Oacute" : 211,
//     "Ocirc" : 212,
//     "Ograve" : 210,
//     "Oslash" : 216,
//     "Otilde" : 213,
//     "Ouml" : 214,
//     "THORN" : 222,
//     "Uacute" : 218,
//     "Ucirc" : 219,
//     "Ugrave" : 217,
//     "Uuml" : 220,
//     "Yacute" : 221,
//     "aacute" : 225,
//     "acirc" : 226,
//     "aelig" : 230,
//     "agrave" : 224,
//     "aring" : 229,
//     "atilde" : 227,
//     "auml" : 228,
//     "ccedil" : 231,
//     "eacute" : 233,
//     "ecirc" : 234,
//     "egrave" : 232,
//     "eth" : 240,
//     "euml" : 235,
//     "iacute" : 237,
//     "icirc" : 238,
//     "igrave" : 236,
//     "iuml" : 239,
//     "ntilde" : 241,
//     "oacute" : 243,
//     "ocirc" : 244,
//     "ograve" : 242,
//     "oslash" : 248,
//     "otilde" : 245,
//     "ouml" : 246,
//     "szlig" : 223,
//     "thorn" : 254,
//     "uacute" : 250,
//     "ucirc" : 251,
//     "ugrave" : 249,
//     "uuml" : 252,
//     "yacute" : 253,
//     "yuml" : 255,
//     "copy" : 169,
//     "reg" : 174,
//     "nbsp" : 160,
//     "iexcl" : 161,
//     "cent" : 162,
//     "pound" : 163,
//     "curren" : 164,
//     "yen" : 165,
//     "brvbar" : 166,
//     "sect" : 167,
//     "uml" : 168,
//     "ordf" : 170,
//     "laquo" : 171,
//     "not" : 172,
//     "shy" : 173,
//     "macr" : 175,
//     "deg" : 176,
//     "plusmn" : 177,
//     "sup1" : 185,
//     "sup2" : 178,
//     "sup3" : 179,
//     "acute" : 180,
//     "micro" : 181,
//     "para" : 182,
//     "middot" : 183,
//     "cedil" : 184,
//     "ordm" : 186,
//     "raquo" : 187,
//     "frac14" : 188,
//     "frac12" : 189,
//     "frac34" : 190,
//     "iquest" : 191,
//     "times" : 215,
//     "divide" : 247,
//     "OElig;" : 338,
//     "oelig;" : 339,
//     "Scaron;" : 352,
//     "scaron;" : 353,
//     "Yuml;" : 376,
//     "fnof;" : 402,
//     "circ;" : 710,
//     "tilde;" : 732,
//     "Alpha;" : 913,
//     "Beta;" : 914,
//     "Gamma;" : 915,
//     "Delta;" : 916,
//     "Epsilon;" : 917,
//     "Zeta;" : 918,
//     "Eta;" : 919,
//     "Theta;" : 920,
//     "Iota;" : 921,
//     "Kappa;" : 922,
//     "Lambda;" : 923,
//     "Mu;" : 924,
//     "Nu;" : 925,
//     "Xi;" : 926,
//     "Omicron;" : 927,
//     "Pi;" : 928,
//     "Rho;" : 929,
//     "Sigma;" : 931,
//     "Tau;" : 932,
//     "Upsilon;" : 933,
//     "Phi;" : 934,
//     "Chi;" : 935,
//     "Psi;" : 936,
//     "Omega;" : 937,
//     "alpha;" : 945,
//     "beta;" : 946,
//     "gamma;" : 947,
//     "delta;" : 948,
//     "epsilon;" : 949,
//     "zeta;" : 950,
//     "eta;" : 951,
//     "theta;" : 952,
//     "iota;" : 953,
//     "kappa;" : 954,
//     "lambda;" : 955,
//     "mu;" : 956,
//     "nu;" : 957,
//     "xi;" : 958,
//     "omicron;" : 959,
//     "pi;" : 960,
//     "rho;" : 961,
//     "sigmaf;" : 962,
//     "sigma;" : 963,
//     "tau;" : 964,
//     "upsilon;" : 965,
//     "phi;" : 966,
//     "chi;" : 967,
//     "psi;" : 968,
//     "omega;" : 969,
//     "thetasym;" : 977,
//     "upsih;" : 978,
//     "piv;" : 982,
//     "ensp;" : 8194,
//     "emsp;" : 8195,
//     "thinsp;" : 8201,
//     "zwnj;" : 8204,
//     "zwj;" : 8205,
//     "lrm;" : 8206,
//     "rlm;" : 8207,
//     "ndash;" : 8211,
//     "mdash;" : 8212,
//     "lsquo;" : 8216,
//     "rsquo;" : 8217,
//     "sbquo;" : 8218,
//     "ldquo;" : 8220,
//     "rdquo;" : 8221,
//     "bdquo;" : 8222,
//     "dagger;" : 8224,
//     "Dagger;" : 8225,
//     "bull;" : 8226,
//     "hellip;" : 8230,
//     "permil;" : 8240,
//     "prime;" : 8242,
//     "Prime;" : 8243,
//     "lsaquo;" : 8249,
//     "rsaquo;" : 8250,
//     "oline;" : 8254,
//     "frasl;" : 8260,
//     "euro;" : 8364,
//     "image;" : 8465,
//     "weierp;" : 8472,
//     "real;" : 8476,
//     "trade;" : 8482,
//     "alefsym;" : 8501,
//     "larr;" : 8592,
//     "uarr;" : 8593,
//     "rarr;" : 8594,
//     "darr;" : 8595,
//     "harr;" : 8596,
//     "crarr;" : 8629,
//     "lArr;" : 8656,
//     "uArr;" : 8657,
//     "rArr;" : 8658,
//     "dArr;" : 8659,
//     "hArr;" : 8660,
//     "forall;" : 8704,
//     "part;" : 8706,
//     "exist;" : 8707,
//     "empty;" : 8709,
//     "nabla;" : 8711,
//     "isin;" : 8712,
//     "notin;" : 8713,
//     "ni;" : 8715,
//     "prod;" : 8719,
//     "sum;" : 8721,
//     "minus;" : 8722,
//     "lowast;" : 8727,
//     "radic;" : 8730,
//     "prop;" : 8733,
//     "infin;" : 8734,
//     "ang;" : 8736,
//     "and;" : 8743,
//     "or;" : 8744,
//     "cap;" : 8745,
//     "cup;" : 8746,
//     "int;" : 8747,
//     "there4;" : 8756,
//     "sim;" : 8764,
//     "cong;" : 8773,
//     "asymp;" : 8776,
//     "ne;" : 8800,
//     "equiv;" : 8801,
//     "le;" : 8804,
//     "ge;" : 8805,
//     "sub;" : 8834,
//     "sup;" : 8835,
//     "nsub;" : 8836,
//     "sube;" : 8838,
//     "supe;" : 8839,
//     "oplus;" : 8853,
//     "otimes;" : 8855,
//     "perp;" : 8869,
//     "sdot;" : 8901,
//     "lceil;" : 8968,
//     "rceil;" : 8969,
//     "lfloor;" : 8970,
//     "rfloor;" : 8971,
//     "lang;" : 9001,
//     "rang;" : 9002,
//     "loz;" : 9674,
//     "spades;" : 9824,
//     "clubs;" : 9827,
//     "hearts;" : 9829,
//     "diams;" : 9830
//   }
//
//
// }).call(this);
