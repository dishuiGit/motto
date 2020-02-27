package lunr

import (
	"github.com/ddliu/motto"
	"github.com/robertkrimen/otto"
)

func lunr(vm *motto.Motto) (otto.Value, error) {
	return motto.CreateLoaderFromSource(source(), "")(vm)
}

func source() string {
	return `var e=function(t){var n=new e.Index;return n.pipeline.add(e.trimmer,e.stopWordFilter,e.stemmer),t&&t.call(n,n),n};e.version="0.7.1",e.utils={},e.utils.warn=function(e){return function(t){e.console&&console.warn&&console.warn(t)}}(this),e.utils.asString=function(e){return void 0===e||null===e?"":e.toString()},e.EventEmitter=function(){this.events={}},e.EventEmitter.prototype.addListener=function(){var e=Array.prototype.slice.call(arguments),t=e.pop(),n=e;if("function"!=typeof t)throw new TypeError("last argument must be a function");n.forEach(function(e){this.hasHandler(e)||(this.events[e]=[]),this.events[e].push(t)},this)},e.EventEmitter.prototype.removeListener=function(e,t){if(this.hasHandler(e)){var n=this.events[e].indexOf(t);this.events[e].splice(n,1),this.events[e].length||delete this.events[e]}},e.EventEmitter.prototype.emit=function(e){if(this.hasHandler(e)){var t=Array.prototype.slice.call(arguments,1);this.events[e].forEach(function(e){e.apply(void 0,t)})}},e.EventEmitter.prototype.hasHandler=function(e){return e in this.events},e.tokenizer=function(t){if(!arguments.length||null==t||void 0==t)return[];if(Array.isArray(t))return t.map(function(t){return e.utils.asString(t).toLowerCase()});var n=t.toString().trim().toLowerCase();return nodejieba.cut(n)},e.tokenizer.seperator=/[\s\-]+/,e.tokenizer.load=function(e){var t=this.registeredFunctions[e];if(!t)throw new Error("Cannot load un-registered function: "+e);return t},e.tokenizer.label="default",e.tokenizer.registeredFunctions={"default":e.tokenizer},e.tokenizer.registerFunction=function(t,n){n in this.registeredFunctions&&e.utils.warn("Overwriting existing tokenizer: "+n),t.label=n,this.registeredFunctions[n]=t},e.Pipeline=function(){this._stack=[]},e.Pipeline.registeredFunctions={},e.Pipeline.registerFunction=function(t,n){n in this.registeredFunctions&&e.utils.warn("Overwriting existing registered function: "+n),t.label=n,e.Pipeline.registeredFunctions[t.label]=t},e.Pipeline.warnIfFunctionNotRegistered=function(t){var n=t.label&&t.label in this.registeredFunctions;n||e.utils.warn("Function is not registered with pipeline. This may cause problems when serialising the index.\n",t)},e.Pipeline.load=function(t){var n=new e.Pipeline;return t.forEach(function(t){var i=e.Pipeline.registeredFunctions[t];if(!i)throw new Error("Cannot load un-registered function: "+t);n.add(i)}),n},e.Pipeline.prototype.add=function(){var t=Array.prototype.slice.call(arguments);t.forEach(function(t){e.Pipeline.warnIfFunctionNotRegistered(t),this._stack.push(t)},this)},e.Pipeline.prototype.after=function(t,n){e.Pipeline.warnIfFunctionNotRegistered(n);var i=this._stack.indexOf(t);if(-1==i)throw new Error("Cannot find existingFn");i+=1,this._stack.splice(i,0,n)},e.Pipeline.prototype.before=function(t,n){e.Pipeline.warnIfFunctionNotRegistered(n);var i=this._stack.indexOf(t);if(-1==i)throw new Error("Cannot find existingFn");this._stack.splice(i,0,n)},e.Pipeline.prototype.remove=function(e){var t=this._stack.indexOf(e);-1!=t&&this._stack.splice(t,1)},e.Pipeline.prototype.run=function(e){for(var t=[],n=e.length,i=this._stack.length,r=0;n>r;r++){for(var o=e[r],s=0;i>s&&(o=this._stack[s](o,r,e),void 0!==o&&""!==o);s++);void 0!==o&&""!==o&&t.push(o)}return t},e.Pipeline.prototype.reset=function(){this._stack=[]},e.Pipeline.prototype.toJSON=function(){return this._stack.map(function(t){return e.Pipeline.warnIfFunctionNotRegistered(t),t.label})},e.Vector=function(){this._magnitude=null,this.list=void 0,this.length=0},e.Vector.Node=function(e,t,n){this.idx=e,this.val=t,this.next=n},e.Vector.prototype.insert=function(t,n){this._magnitude=void 0;var i=this.list;if(!i)return this.list=new e.Vector.Node(t,n,i),this.length++;if(t<i.idx)return this.list=new e.Vector.Node(t,n,i),this.length++;for(var r=i,o=i.next;void 0!=o;){if(t<o.idx)return r.next=new e.Vector.Node(t,n,o),this.length++;r=o,o=o.next}return r.next=new e.Vector.Node(t,n,o),this.length++},e.Vector.prototype.magnitude=function(){if(this._magnitude)return this._magnitude;for(var e,t=this.list,n=0;t;)e=t.val,n+=e*e,t=t.next;return this._magnitude=Math.sqrt(n)},e.Vector.prototype.dot=function(e){for(var t=this.list,n=e.list,i=0;t&&n;)t.idx<n.idx?t=t.next:t.idx>n.idx?n=n.next:(i+=t.val*n.val,t=t.next,n=n.next);return i},e.Vector.prototype.similarity=function(e){return this.dot(e)/(this.magnitude()*e.magnitude())},e.SortedSet=function(){this.length=0,this.elements=[]},e.SortedSet.load=function(e){var t=new this;return t.elements=e,t.length=e.length,t},e.SortedSet.prototype.add=function(){var e,t;for(e=0;e<arguments.length;e++)t=arguments[e],~this.indexOf(t)||this.elements.splice(this.locationFor(t),0,t);this.length=this.elements.length},e.SortedSet.prototype.toArray=function(){return this.elements.slice()},e.SortedSet.prototype.map=function(e,t){return this.elements.map(e,t)},e.SortedSet.prototype.forEach=function(e,t){return this.elements.forEach(e,t)},e.SortedSet.prototype.indexOf=function(e){for(var t=0,n=this.elements.length,i=n-t,r=t+Math.floor(i/2),o=this.elements[r];i>1;){if(o===e)return r;e>o&&(t=r),o>e&&(n=r),i=n-t,r=t+Math.floor(i/2),o=this.elements[r]}return o===e?r:-1},e.SortedSet.prototype.locationFor=function(e){for(var t=0,n=this.elements.length,i=n-t,r=t+Math.floor(i/2),o=this.elements[r];i>1;)e>o&&(t=r),o>e&&(n=r),i=n-t,r=t+Math.floor(i/2),o=this.elements[r];return o>e?r:e>o?r+1:void 0},e.SortedSet.prototype.intersect=function(t){for(var n=new e.SortedSet,i=0,r=0,o=this.length,s=t.length,a=this.elements,h=t.elements;;){if(i>o-1||r>s-1)break;a[i]!==h[r]?a[i]<h[r]?i++:a[i]>h[r]&&r++:(n.add(a[i]),i++,r++)}return n},e.SortedSet.prototype.clone=function(){var t=new e.SortedSet;return t.elements=this.toArray(),t.length=t.elements.length,t},e.SortedSet.prototype.union=function(e){var t,n,i;this.length>=e.length?(t=this,n=e):(t=e,n=this),i=t.clone();for(var r=0,o=n.toArray();r<o.length;r++)i.add(o[r]);return i},e.SortedSet.prototype.toJSON=function(){return this.toArray()},e.Index=function(){this._fields=[],this._ref="id",this.pipeline=new e.Pipeline,this.documentStore=new e.Store,this.tokenStore=new e.TokenStore,this.corpusTokens=new e.SortedSet,this.eventEmitter=new e.EventEmitter,this.tokenizerFn=e.tokenizer,this._idfCache={},this.on("add","remove","update",function(){this._idfCache={}}.bind(this))},e.Index.prototype.on=function(){var e=Array.prototype.slice.call(arguments);return this.eventEmitter.addListener.apply(this.eventEmitter,e)},e.Index.prototype.off=function(e,t){return this.eventEmitter.removeListener(e,t)},e.Index.load=function(t){t.version!==e.version&&e.utils.warn("version mismatch: current "+e.version+" importing "+t.version);var n=new this;return n._fields=t.fields,n._ref=t.ref,n.tokenizer=e.tokenizer.load(t.tokenizer),n.documentStore=e.Store.load(t.documentStore),n.tokenStore=e.TokenStore.load(t.tokenStore),n.corpusTokens=e.SortedSet.load(t.corpusTokens),n.pipeline=e.Pipeline.load(t.pipeline),n},e.Index.prototype.field=function(e,t){var t=t||{},n={name:e,boost:t.boost||1};return this._fields.push(n),this},e.Index.prototype.ref=function(e){return this._ref=e,this},e.Index.prototype.tokenizer=function(t){var n=t.label&&t.label in e.tokenizer.registeredFunctions;return n||e.utils.warn("Function is not a registered tokenizer. This may cause problems when serialising the index"),this.tokenizerFn=t,this},e.Index.prototype.add=function(t,n){var i={},r=new e.SortedSet,o=t[this._ref],n=void 0===n?!0:n;this._fields.forEach(function(e){var n=this.pipeline.run(this.tokenizerFn(t[e.name]));i[e.name]=n;for(var o=0;o<n.length;o++){var s=n[o];r.add(s),this.corpusTokens.add(s)}},this),this.documentStore.set(o,r);for(var s=0;s<r.length;s++){for(var a=r.elements[s],h=0,u=0;u<this._fields.length;u++){var l=this._fields[u],c=i[l.name],d=c.length;if(d){for(var f=0,p=0;d>p;p++)c[p]===a&&f++;h+=f/d*l.boost}}this.tokenStore.add(a,{ref:o,tf:h})}n&&this.eventEmitter.emit("add",t,this)},e.Index.prototype.remove=function(e,t){var n=e[this._ref],t=void 0===t?!0:t;if(this.documentStore.has(n)){var i=this.documentStore.get(n);this.documentStore.remove(n),i.forEach(function(e){this.tokenStore.remove(e,n)},this),t&&this.eventEmitter.emit("remove",e,this)}},e.Index.prototype.update=function(e,t){var t=void 0===t?!0:t;this.remove(e,!1),this.add(e,!1),t&&this.eventEmitter.emit("update",e,this)},e.Index.prototype.idf=function(e){var t="@"+e;if(Object.prototype.hasOwnProperty.call(this._idfCache,t))return this._idfCache[t];var n=this.tokenStore.count(e),i=1;return n>0&&(i=1+Math.log(this.documentStore.length/n)),this._idfCache[t]=i},e.Index.prototype.search=function(t){var n=this.pipeline.run(this.tokenizerFn(t)),i=new e.Vector,r=[],o=this._fields.reduce(function(e,t){return e+t.boost},0),s=n.some(function(e){return this.tokenStore.has(e)},this);if(!s)return[];n.forEach(function(t,n,s){var a=1/s.length*this._fields.length*o,h=this,u=this.tokenStore.expand(t).reduce(function(n,r){var o=h.corpusTokens.indexOf(r),s=h.idf(r),u=1,l=new e.SortedSet;if(r!==t){var c=Math.max(3,r.length-t.length);u=1/Math.log(c)}o>-1&&i.insert(o,a*s*u);for(var d=h.tokenStore.get(r),f=Object.keys(d),p=f.length,v=0;p>v;v++)l.add(d[f[v]].ref);return n.union(l)},new e.SortedSet);r.push(u)},this);var a=r.reduce(function(e,t){return e.intersect(t)});return a.map(function(e){return{ref:e,score:i.similarity(this.documentVector(e))}},this).sort(function(e,t){return t.score-e.score})},e.Index.prototype.documentVector=function(t){for(var n=this.documentStore.get(t),i=n.length,r=new e.Vector,o=0;i>o;o++){var s=n.elements[o],a=this.tokenStore.get(s)[t].tf,h=this.idf(s);r.insert(this.corpusTokens.indexOf(s),a*h)}return r},e.Index.prototype.toJSON=function(){return{version:e.version,fields:this._fields,ref:this._ref,tokenizer:this.tokenizerFn.label,documentStore:this.documentStore.toJSON(),tokenStore:this.tokenStore.toJSON(),corpusTokens:this.corpusTokens.toJSON(),pipeline:this.pipeline.toJSON()}},e.Index.prototype.use=function(e){var t=Array.prototype.slice.call(arguments,1);t.unshift(this),e.apply(this,t)},e.Store=function(){this.store={},this.length=0},e.Store.load=function(t){var n=new this;return n.length=t.length,n.store=Object.keys(t.store).reduce(function(n,i){return n[i]=e.SortedSet.load(t.store[i]),n},{}),n},e.Store.prototype.set=function(e,t){this.has(e)||this.length++,this.store[e]=t},e.Store.prototype.get=function(e){return this.store[e]},e.Store.prototype.has=function(e){return e in this.store},e.Store.prototype.remove=function(e){this.has(e)&&(delete this.store[e],this.length--)},e.Store.prototype.toJSON=function(){return{store:this.store,length:this.length}},e.stemmer=function(){var e={ational:"ate",tional:"tion",enci:"ence",anci:"ance",izer:"ize",bli:"ble",alli:"al",entli:"ent",eli:"e",ousli:"ous",ization:"ize",ation:"ate",ator:"ate",alism:"al",iveness:"ive",fulness:"ful",ousness:"ous",aliti:"al",iviti:"ive",biliti:"ble",logi:"log"},t={icate:"ic",ative:"",alize:"al",iciti:"ic",ical:"ic",ful:"",ness:""},n="[^aeiou]",i="[aeiouy]",r=n+"[^aeiouy]*",o=i+"[aeiou]*",s="^("+r+")?"+o+r,a="^("+r+")?"+o+r+"("+o+")?$",h="^("+r+")?"+o+r+o+r,u="^("+r+")?"+i,l=new RegExp(s),c=new RegExp(h),d=new RegExp(a),f=new RegExp(u),p=/^(.+?)(ss|i)es$/,v=/^(.+?)([^s])s$/,g=/^(.+?)eed$/,m=/^(.+?)(ed|ing)$/,y=/.$/,S=/(at|bl|iz)$/,w=new RegExp("([^aeiouylsz])$"),k=new RegExp("^"+r+i+"[^aeiouwxy]$"),x=/^(.+?[^aeiou])y$/,b=/^(.+?)(ational|tional|enci|anci|izer|bli|alli|entli|eli|ousli|ization|ation|ator|alism|iveness|fulness|ousness|aliti|iviti|biliti|logi)$/,E=/^(.+?)(icate|ative|alize|iciti|ical|ful|ness)$/,_=/^(.+?)(al|ance|ence|er|ic|able|ible|ant|ement|ment|ent|ou|ism|ate|iti|ous|ive|ize)$/,F=/^(.+?)(s|t)(ion)$/,z=/^(.+?)e$/,O=/ll$/,P=new RegExp("^"+r+i+"[^aeiouwxy]$"),T=function(n){var i,r,o,s,a,h,u;if(n.length<3)return n;if(o=n.substr(0,1),"y"==o&&(n=o.toUpperCase()+n.substr(1)),s=p,a=v,s.test(n)?n=n.replace(s,"$1$2"):a.test(n)&&(n=n.replace(a,"$1$2")),s=g,a=m,s.test(n)){var T=s.exec(n);s=l,s.test(T[1])&&(s=y,n=n.replace(s,""))}else if(a.test(n)){var T=a.exec(n);i=T[1],a=f,a.test(i)&&(n=i,a=S,h=w,u=k,a.test(n)?n+="e":h.test(n)?(s=y,n=n.replace(s,"")):u.test(n)&&(n+="e"))}if(s=x,s.test(n)){var T=s.exec(n);i=T[1],n=i+"i"}if(s=b,s.test(n)){var T=s.exec(n);i=T[1],r=T[2],s=l,s.test(i)&&(n=i+e[r])}if(s=E,s.test(n)){var T=s.exec(n);i=T[1],r=T[2],s=l,s.test(i)&&(n=i+t[r])}if(s=_,a=F,s.test(n)){var T=s.exec(n);i=T[1],s=c,s.test(i)&&(n=i)}else if(a.test(n)){var T=a.exec(n);i=T[1]+T[2],a=c,a.test(i)&&(n=i)}if(s=z,s.test(n)){var T=s.exec(n);i=T[1],s=c,a=d,h=P,(s.test(i)||a.test(i)&&!h.test(i))&&(n=i)}return s=O,a=c,s.test(n)&&a.test(n)&&(s=y,n=n.replace(s,"")),"y"==o&&(n=o.toLowerCase()+n.substr(1)),n};return T}(),e.Pipeline.registerFunction(e.stemmer,"stemmer"),e.generateStopWordFilter=function(e){var t=e.reduce(function(e,t){return e[t]=t,e},{});return function(e){return e&&t[e]!==e?e:void 0}},e.stopWordFilter=e.generateStopWordFilter(["a","able","about","across","after","all","almost","also","am","among","an","and","any","are","as","at","be","because","been","but","by","can","cannot","could","dear","did","do","does","either","else","ever","every","for","from","get","got","had","has","have","he","her","hers","him","his","how","however","i","if","in","into","is","it","its","just","least","let","like","likely","may","me","might","most","must","my","neither","no","nor","not","of","off","often","on","only","or","other","our","own","rather","said","say","says","she","should","since","so","some","than","that","the","their","them","then","there","these","they","this","tis","to","too","twas","us","wants","was","we","were","what","when","where","which","while","who","whom","why","will","with","would","yet","you","your"]),e.Pipeline.registerFunction(e.stopWordFilter,"stopWordFilter"),e.trimmer=function(e){var t=e.replace(/^\s+/,"").replace(/\s+$/,"");return""===t?void 0:t},e.Pipeline.registerFunction(e.trimmer,"trimmer"),e.TokenStore=function(){this.root={docs:{}},this.length=0},e.TokenStore.load=function(e){var t=new this;return t.root=e.root,t.length=e.length,t},e.TokenStore.prototype.add=function(e,t,n){var n=n||this.root,i=e.charAt(0),r=e.split(""),r=r.slice(1,r.length).join("");return i in n||(n[i]={docs:{}}),0===r.length?(n[i].docs[t.ref]=t,void(this.length+=1)):this.add(r,t,n[i])},e.TokenStore.prototype.has=function(e){if(!e)return!1;for(var t=this.root,n=0;n<e.length;n++){if(!t[e.charAt(n)])return!1;t=t[e.charAt(n)]}return!0},e.TokenStore.prototype.getNode=function(e){if(!e)return{};for(var t=this.root,n=0;n<e.length;n++){if(!t[e.charAt(n)])return{};t=t[e.charAt(n)]}return t},e.TokenStore.prototype.get=function(e,t){return this.getNode(e,t).docs||{}},e.TokenStore.prototype.count=function(e,t){return Object.keys(this.get(e,t)).length},e.TokenStore.prototype.remove=function(e,t){if(e){for(var n=this.root,i=0;i<e.length;i++){if(!(e.charAt(i)in n))return;n=n[e.charAt(i)]}delete n.docs[t]}},e.TokenStore.prototype.expand=function(e,t){var n=this.getNode(e),i=n.docs||{},t=t||[];return Object.keys(i).length&&t.push(e),Object.keys(n).forEach(function(n){"docs"!==n&&t.concat(this.expand(e+n,t))},this),t},e.TokenStore.prototype.toJSON=function(){return{root:this.root,length:this.length}},function(e,t){"function"==typeof define&&define.amd?define(t):"object"==typeof exports?module.exports=t():e.lunr=t()}(this,function(){return e});`
}


func init() {
	motto.AddModule("lunr", lunr)
}