# PHP functions implemented by Golang

[![GoDoc](https://godoc.org/github.com/imartingraham/php?status.svg)](https://godoc.org/github.com/imartingraham/php)
[![Build Status](https://travis-ci.org/imartingraham/php.svg?branch=master)](https://travis-ci.org/imartingraham/php)
[![](https://goreportcard.com/badge/github.com/imartingraham/php)](https://goreportcard.com/report/github.com/imartingraham/php)
[![codecov](https://codecov.io/gh/imartingraham/php/branch/master/graph/badge.svg)](https://codecov.io/gh/imartingraham/php)
[![Release](https://img.shields.io/github/release/imartingraham/php.svg)](https://github.com/imartingraham/php/releases)

This package implements some PHP functions by Golang. Please note that it's impossible to have 100% the same behaviour between PHP and Golang functions because their mechanisms are quite different.

Minimum go version requirement:

| OS      | Go version |
| ------- | ---------- |
| Linux   | 1.9        |
| OSX     | 1.12       |
| Windows | 1.13       |


### Date/Time Functions

| PHP function                             | Golang function                          |
| ---------------------------------------- | ---------------------------------------- |
| [checkdate](https://php.net/manual/en/function.checkdate.php) | [Checkdate](https://godoc.org/github.com/imartingraham/php#Checkdate) |
| [date](https://php.net/manual/en/function.date.php) | [Date](https://godoc.org/github.com/imartingraham/php#Date) |
| [strtotime](https://php.net/manual/en/function.strtotime.php) | [Strtotime](https://godoc.org/github.com/imartingraham/php#Strtotime) |
| [time](https://php.net/manual/en/function.time.php) | [Time](https://godoc.org/github.com/imartingraham/php#Time) |
| [microtime](https://php.net/manual/en/function.microtime.php) | [Microtime](https://godoc.org/github.com/imartingraham/php#Microtime) |
| [date_add](https://www.php.net/manual/en/function.date-add.php) | [DateAdd](https://godoc.org/github.com/imartingraham/php#DateAdd) |
| [date_create_from_format](https://www.php.net/manual/en/function.date-create-from-format.php) | [DateCreateFromFormat](https://godoc.org/github.com/imartingraham/php#DateCreateFromFormat) |
| [date_create](https://www.php.net/manual/en/function.date-create.php) | [DateCreate](https://godoc.org/github.com/imartingraham/php#DateCreate) |
| [date_date_set](https://www.php.net/manual/en/function.date-date-set.php) | [DateDateSet](https://godoc.org/github.com/imartingraham/php#DateDateSet) |
| [date_default_timezone_get](https://www.php.net/manual/en/function.date-default-timezone-get.php) | [DateDefaultTimezoneGet](https://godoc.org/github.com/imartingraham/php#DateDefaultTimezoneGet) |
| [date_default_timezone_set](https://www.php.net/manual/en/function.date-default-timezone-set.php) | [DateDefaultTimezoneSet](https://godoc.org/github.com/imartingraham/php#DateDefaultTimezoneSet) |
| [date_diff](https://www.php.net/manual/en/function.date-diff.php) | [DateDiff](https://godoc.org/github.com/imartingraham/php#DateDiff) |
| [date_format](https://www.php.net/manual/en/function.date-format.php) | [DateFormat](https://godoc.org/github.com/imartingraham/php#DateFormat) |
| [date_interval_create_from_date_string](https://www.php.net/manual/en/function.date-interval-create-from-date-string.php) | [DateIntervalCreateFromDateString](https://godoc.org/github.com/imartingraham/php#DateIntervalCreateFromDateString) |
| [date_isodate_set](https://www.php.net/manual/en/function.date-isodate-set.php) | [DateISODateSet](https://godoc.org/github.com/imartingraham/php#DateISODateSet) |
| [date_modify](https://www.php.net/manual/en/function.date-modify.php) | [DateModify](https://godoc.org/github.com/imartingraham/php#DateModify) |
| [date_offset_get](https://www.php.net/manual/en/function.date-offset-get.php) | [DateOffsetGet](https://godoc.org/github.com/imartingraham/php#DateOffsetGet) |
| [date_sub](https://www.php.net/manual/en/function.date-sub.php) | [DateSub](https://godoc.org/github.com/imartingraham/php#DateSub) |
| [date_timestamp_get](https://www.php.net/manual/en/function.date-timestamp-get.php) | [DateTimestampGet](https://godoc.org/github.com/imartingraham/php#DateTimestampGet) |
| [date_timestamp_set](https://www.php.net/manual/en/function.date-timestamp-set.php) | [DateTimestampSet](https://godoc.org/github.com/imartingraham/php#DateTimestampSet) |
| [date_timezone_get](https://www.php.net/manual/en/function.date-timezone-get.php) | [DateTimezoneGet](https://godoc.org/github.com/imartingraham/php#DateTimezoneGet) |
| [date_timezone_set](https://www.php.net/manual/en/function.date-timezone-set.php) | [DateTimezoneSet](https://godoc.org/github.com/imartingraham/php#DateTimezoneSet) |

### String Functions

| PHP function                             | Golang function                          |
| ---------------------------------------- | ---------------------------------------- |
| [addslashes](https://php.net/manual/en/function.addslashes.php) | [Addslashes](https://godoc.org/github.com/imartingraham/php#Addslashes) |
| [chr](https://php.net/manual/en/function.chr.php) | [Chr](https://godoc.org/github.com/imartingraham/php#Chr) |
| [crc32](https://php.net/manual/en/function.crc32.php) | [Crc32](https://godoc.org/github.com/imartingraham/php#Crc32) |
| [explode](https://php.net/manual/en/function.explode.php) | [Explode](https://godoc.org/github.com/imartingraham/php#Explode) |
| [bin2hex](https://php.net/manual/en/function.bin2hex.php) | [Bin2hex](https://godoc.org/github.com/imartingraham/php#Bin2hex) |
| [hex2bin](https://php.net/manual/en/function.hex2bin.php) | [Hex2bin](https://godoc.org/github.com/imartingraham/php#Hex2bin) |
| [htmlspecialchars](https://php.net/manual/en/function.htmlspecialchars.php) | [HTMLSpecialchars](https://godoc.org/github.com/imartingraham/php#HTMLSpecialchars) |
| [htmlspecialchars_decode](https://php.net/manual/en/function.htmlspecialchars-decode.php) | [HTMLSpecialcharsDecode](https://godoc.org/github.com/imartingraham/php#HTMLSpecialcharsDecode) |
| [implode](https://php.net/manual/en/function.implode.php) | [Implode](https://godoc.org/github.com/imartingraham/php#Implode) |
| [str_ireplace](https://php.net/manual/en/function.str-ireplace.php) | [Ireplace](https://godoc.org/github.com/imartingraham/php#Ireplace) |
| [lcfirst](https://php.net/manual/en/function.lcfirst.php) | [Lcfirst](https://godoc.org/github.com/imartingraham/php#Lcfirst) |
| [md5](https://php.net/manual/en/function.md5.php) | [Md5](https://godoc.org/github.com/imartingraham/php#Md5) |
| [md5_file](https://php.net/manual/en/function.md5-file.php) | [Md5File](https://godoc.org/github.com/imartingraham/php#Md5File) |
| [sha1](https://php.net/manual/en/function.sha1.php) | [Sha1](https://godoc.org/github.com/imartingraham/php#Sha1) |
| [sha1_file](https://php.net/manual/en/function.sha1-file.php) | [Sha1File](https://godoc.org/github.com/imartingraham/php#Sha1File) |
| [number_format](https://php.net/manual/en/function.number-format.php) | [NumberFormat](https://godoc.org/github.com/imartingraham/php#NumberFormat) |
| [ord](https://php.net/manual/en/function.ord.php) | [Ord](https://godoc.org/github.com/imartingraham/php#Ord) |
| [str_replace](https://php.net/manual/en/function.str-replace.php) | [Replace](https://godoc.org/github.com/imartingraham/php#Replace) |
| [similar_text](https://php.net/manual/en/function.similar-text.php) | [SimilarText](https://godoc.org/github.com/imartingraham/php#SimilarText) |
| [stripslashes](https://php.net/manual/en/function.stripslashes.php) | [Stripslashes](https://godoc.org/github.com/imartingraham/php#Stripslashes) |
| [mb_stripos](https://php.net/manual/en/function.mb-stripos.php) | [Stripos](https://godoc.org/github.com/imartingraham/php#Stripos) |
| [mb_stristr](https://php.net/manual/en/function.mb-stristr.php) | [Stristr](https://godoc.org/github.com/imartingraham/php#Stristr) |
| [mb_strlen](https://php.net/manual/en/function.mb-strlen.php) | [Strlen](https://godoc.org/github.com/imartingraham/php#Strlen) |
| [mb_strpos](https://php.net/manual/en/function.mb-strpos.php) | [Strpos](https://godoc.org/github.com/imartingraham/php#Strpos) |
| [mb_strripos](https://php.net/manual/en/function.mb-strripos.php) | [Strripos](https://godoc.org/github.com/imartingraham/php#Strripos) |
| [mb_strrpos](https://php.net/manual/en/function.mb-strrpos.php) | [Strrpos](https://godoc.org/github.com/imartingraham/php#Strrpos) |
| [mb_strstr](https://php.net/manual/en/function.mb-strstr.php) | [Strstr](https://godoc.org/github.com/imartingraham/php#Strstr) |
| [mb_substr](https://php.net/manual/en/function.mb-substr.php) | [Substr](https://godoc.org/github.com/imartingraham/php#Substr) |
| [str_pad](https://php.net/manual/en/function.str-pad.php) | [StrPad](https://godoc.org/github.com/imartingraham/php#StrPad) |
| [str_repeat](https://php.net/manual/en/function.str-repeat.php) | [StrRepeat](https://godoc.org/github.com/imartingraham/php#StrRepeat) |
| [strrev](https://php.net/manual/en/function.strrev.php) | [Strrev](https://godoc.org/github.com/imartingraham/php#Strrev) |
| [strtolower](https://php.net/manual/en/function.strtolower.php) | [Strtolower](https://godoc.org/github.com/imartingraham/php#Strtolower) |
| [strtoupper](https://php.net/manual/en/function.strtoupper.php) | [Strtoupper](https://godoc.org/github.com/imartingraham/php#Strtoupper) |
| [str_word_count](https://php.net/manual/en/function.str-word-count.php) | [StrWordCount](https://godoc.org/github.com/imartingraham/php#StrWordCount) |
| [ltrim](https://php.net/manual/en/function.ltrim.php) | [Ltrim](https://godoc.org/github.com/imartingraham/php#Ltrim) |
| [rtrim](https://php.net/manual/en/function.rtrim.php) | [Rtrim](https://godoc.org/github.com/imartingraham/php#Rtrim) |
| [trim](https://php.net/manual/en/function.trim.php) | [Trim](https://godoc.org/github.com/imartingraham/php#Trim) |
| [ucfirst](https://php.net/manual/en/function.ucfirst.php) | [Ucfirst](https://godoc.org/github.com/imartingraham/php#Ucfirst) |
| [ucwords](https://php.net/manual/en/function.ucwords.php) | [Ucwords](https://godoc.org/github.com/imartingraham/php#Ucwords) |

### Math Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [abs](https://php.net/manual/en/function.abs.php) | [Abs](https://godoc.org/github.com/imartingraham/php#Abs) |
| [base_convert](https://php.net/manual/en/function.base-convert.php) | [BaseConvert](https://godoc.org/github.com/imartingraham/php#BaseConvert) |
| [bindec](https://php.net/manual/en/function.bindec.php) | [Bindec](https://godoc.org/github.com/imartingraham/php#Bindec) |
| [decbin](https://php.net/manual/en/function.decbin.php) | [Decbin](https://godoc.org/github.com/imartingraham/php#Decbin) |
| [dechex](https://php.net/manual/en/function.dechex.php) | [Dechex](https://godoc.org/github.com/imartingraham/php#Dechex) |
| [hexdec](https://php.net/manual/en/function.hexdec.php) | [Hexdec](https://godoc.org/github.com/imartingraham/php#Hexdec) |
| [decoct](https://php.net/manual/en/function.decoct.php) | [Decoct](https://godoc.org/github.com/imartingraham/php#Decoct) |
| [octdec](https://php.net/manual/en/function.octdec.php) | [Octdec](https://godoc.org/github.com/imartingraham/php#Octdec) |
| [ceil](https://php.net/manual/en/function.ceil.php) | [Ceil](https://godoc.org/github.com/imartingraham/php#Ceil) |
| [floor](https://php.net/manual/en/function.floor.php) | [Floor](https://godoc.org/github.com/imartingraham/php#Floor) |
| [pi](https://php.net/manual/en/function.pi.php) | [Pi](https://godoc.org/github.com/imartingraham/php#Pi) |
| [mt_rand](https://php.net/manual/en/function.mt-rand.php) | [MtRand](https://godoc.org/github.com/imartingraham/php#MtRand) |
| [rand](https://php.net/manual/en/function.rand.php) | [Rand](https://godoc.org/github.com/imartingraham/php#Rand) |
| [round](https://php.net/manual/en/function.round.php) | [Round](https://godoc.org/github.com/imartingraham/php#Round) |

### Array Functions

| PHP function                                                 | Golang function                                              |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [array_chunk](https://php.net/manual/en/function.array-chunk.php) | [ArrayChunk](https://godoc.org/github.com/imartingraham/php#ArrayChunk) |
| [array_column](https://php.net/manual/en/function.array-column.php) | [ArrayColumn](https://godoc.org/github.com/imartingraham/php#ArrayColumn) |
| [array_combine](https://php.net/manual/en/function.array-combine.php) | [ArrayCombine](https://godoc.org/github.com/imartingraham/php#ArrayCombine) |
| [array_diff](https://php.net/manual/en/function.array-diff.php) | [ArrayDiff](https://godoc.org/github.com/imartingraham/php#ArrayDiff) |
| [array_filter](https://php.net/manual/en/function.array-filter.php) | [ArrayFilter](https://godoc.org/github.com/imartingraham/php#ArrayFilter) |
| [array_flip](https://php.net/manual/en/function.array-flip.php) | [ArrayFlip](https://godoc.org/github.com/imartingraham/php#ArrayFlip) |
| [array_intersect](https://php.net/manual/en/function.array-intersect.php) | [ArrayIntersect](https://godoc.org/github.com/imartingraham/php#ArrayIntersect) |
| [array_keys](https://php.net/manual/en/function.array-keys.php) | [ArrayKeys](https://godoc.org/github.com/imartingraham/php#ArrayKeys) |
| [array_key_exists](https://www.php.net/manual/en/function.array-key-exists.php) | [ArrayKeyExists](https://godoc.org/github.com/imartingraham/php#ArrayKeyExists) |
| [array_pad](https://php.net/manual/en/function.array-pad.php) | [ArrayPad](https://godoc.org/github.com/imartingraham/php#ArrayPad) |
| [array_pop](https://php.net/manual/en/function.array-pop.php) | [ArrayPop](https://godoc.org/github.com/imartingraham/php#ArrayPop) |
| [array_push](https://php.net/manual/en/function.array-push.php) | [ArrayPush](https://godoc.org/github.com/imartingraham/php#ArrayPush) |
| [array_reverse](https://php.net/manual/en/function.array-reverse.php) | [ArraySlice](https://godoc.org/github.com/imartingraham/php#ArraySlice) |
| [array_slice](https://php.net/manual/en/function.array-slice.php) | [ArrayReverse](https://godoc.org/github.com/imartingraham/php#ArrayReverse) |
| [array_sum](https://php.net/manual/en/function.array-sum.php) | [ArraySum](https://godoc.org/github.com/imartingraham/php#ArraySum) |
| [array_shift](https://php.net/manual/en/function.array-shift.php) | [ArrayShift](https://godoc.org/github.com/imartingraham/php#ArrayShift) |
| [array_unshift](https://php.net/manual/en/function.array-unshift.php) | [ArrayUnshift](https://godoc.org/github.com/imartingraham/php#ArrayUnshift) |
| [array_unique](https://php.net/manual/en/function.array-unique.php) | [ArrayUnique](https://godoc.org/github.com/imartingraham/php#ArrayUnique) |
| [count](https://php.net/manual/en/function.count.php) | [Count](https://godoc.org/github.com/imartingraham/php#Count) |
| [in_array](https://php.net/manual/en/function.in-array.php) | [InArray](https://godoc.org/github.com/imartingraham/php#InArray) |
| [key_exists](https://www.php.net/manual/en/function.key-exists.php) | [KeyExists](https://godoc.org/github.com/imartingraham/php#KeyExists) |
| [sort](https://php.net/manual/en/function.sort.php) | [Sort](https://godoc.org/github.com/imartingraham/php#Sort) |
| [rsort](https://php.net/manual/en/function.rsort.php) | [Rsort](https://godoc.org/github.com/imartingraham/php#Rsort) |

### Ctype Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [ctype_alnum](https://php.net/manual/en/function.ctype-alnum.php) | [CtypeAlnum](https://godoc.org/github.com/imartingraham/php#CtypeAlnum) |
| [ctype_alpha](https://php.net/manual/en/function.ctype-alpha.php) | [CtypeAlpha](https://godoc.org/github.com/imartingraham/php#CtypeAlpha) |
| [ctype_cntrl](https://php.net/manual/en/function.ctype-cntrl.php) | [CtypeCntrl](https://godoc.org/github.com/imartingraham/php#CtypeCntrl) |
| [ctype_digit](https://php.net/manual/en/function.ctype-digit.php) | [CtypeDigit](https://godoc.org/github.com/imartingraham/php#CtypeDigit) |
| [ctype_graph](https://php.net/manual/en/function.ctype-graph.php) | [CtypeGraph](https://godoc.org/github.com/imartingraham/php#CtypeGraph) |
| [ctype_lower](https://php.net/manual/en/function.ctype-lower.php) | [CtypeLower](https://godoc.org/github.com/imartingraham/php#CtypeLower) |
| [ctype_print](https://php.net/manual/en/function.ctype-print.php) | [CtypePrint](https://godoc.org/github.com/imartingraham/php#CtypePrint) |
| [ctype_punct](https://php.net/manual/en/function.ctype-punct.php) | [CtypePunct](https://godoc.org/github.com/imartingraham/php#CtypePunct) |
| [ctype_space](https://php.net/manual/en/function.ctype-space.php) | [CtypeSpace](https://godoc.org/github.com/imartingraham/php#CtypeSpace) |
| [ctype_upper](https://php.net/manual/en/function.ctype-upper.php) | [CtypeUpper](https://godoc.org/github.com/imartingraham/php#CtypeUpper) |
| [ctype_xdigit](https://php.net/manual/en/function.ctype-xdigit.php) | [CtypeXdigit](https://godoc.org/github.com/imartingraham/php#CtypeXdigit) |

### Filesystem Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [basename](https://www.php.net/manual/en/function.basename.php) | [Basename](https://godoc.org/github.com/imartingraham/php#Basename) |
| [dirname](https://www.php.net/manual/en/function.dirname.php) | [Dirname](https://godoc.org/github.com/imartingraham/php#Dirname) [DirnameWithLevels](https://godoc.org/github.com/imartingraham/php#DirnameWithLevels) |
| [realpath](https://www.php.net/manual/en/function.realpath.php) | [Realpath](https://godoc.org/github.com/imartingraham/php#realpath) |
| [touch](https://www.php.net/manual/en/function.touch.php) | [Touch](https://godoc.org/github.com/imartingraham/php#Touch) |
| [unlink](https://www.php.net/manual/en/function.unlink.php) | [Unlink](https://godoc.org/github.com/imartingraham/php#Unlink) |
| [mkdir](https://www.php.net/manual/en/function.mkdir.php) | [Mkdir](https://godoc.org/github.com/imartingraham/php#Mkdir) |
| [rmdir](https://www.php.net/manual/en/function.rmdir.php) | [Rmdir](https://godoc.org/github.com/imartingraham/php#Rmdir) |
| [symlink](https://www.php.net/manual/en/function.symlink.php) | [Symlink](https://godoc.org/github.com/imartingraham/php#Symlink) |
| [link](https://www.php.net/manual/en/function.link.php) | [Link](https://godoc.org/github.com/imartingraham/php#Link) |
| [chmod](https://www.php.net/manual/en/function.chmod.php) | [Chmod](https://godoc.org/github.com/imartingraham/php#Chmod) |
| [chown](https://www.php.net/manual/en/function.chown.php) | [Chown](https://godoc.org/github.com/imartingraham/php#Chown) |
| [is_file](https://www.php.net/manual/en/function.is-file.php) | [IsFile](https://godoc.org/github.com/imartingraham/php#IsFile) |
| [is_dir](https://www.php.net/manual/en/function.is-dir.php) | [IsDir](https://godoc.org/github.com/imartingraham/php#IsDir) |
| [is_executable](https://www.php.net/manual/en/function.is-executable.php) | [IsExecutable](https://godoc.org/github.com/imartingraham/php#IsExecutable) |
| [is_readable](https://www.php.net/manual/en/function.is-readable.php) | [IsReadable](https://godoc.org/github.com/imartingraham/php#IsReadable) |
| [is_writable](https://www.php.net/manual/en/function.is-writable.php) | [IsWritable](https://godoc.org/github.com/imartingraham/php#IsWritable) |
| [is_link](https://www.php.net/manual/en/function.is-link.php) | [IsLink](https://godoc.org/github.com/imartingraham/php#IsLink) |
| [file_exists](https://www.php.net/manual/en/function.file-exists.php) | [FileExists](https://godoc.org/github.com/imartingraham/php#FileExists) |
| [filesize](https://www.php.net/manual/en/function.filesize.php) | [Filesize](https://godoc.org/github.com/imartingraham/php#Filesize) |
| [copy](https://www.php.net/manual/en/function.copy.php) | [Copy](https://godoc.org/github.com/imartingraham/php#Copy) |
| [rename](https://www.php.net/manual/en/function.rename.php) | [Rename](https://godoc.org/github.com/imartingraham/php#Rename) |
| [file_get_contents](https://www.php.net/manual/en/function.file-get-contents.php) | [FileGetContents](https://godoc.org/github.com/imartingraham/php#FileGetContents) |
| [file_put_contents](https://www.php.net/manual/en/function.file-put-contents.php) | [FilePutContents](https://godoc.org/github.com/imartingraham/php#FilePutContents) |

### Directory Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [getcwd](https://www.php.net/manual/en/function.getcwd.php) | [Getcwd](https://godoc.org/github.com/imartingraham/php#Getcwd) |
| [chdir](https://www.php.net/manual/en/function.chdir.php) | [Chdir](https://godoc.org/github.com/imartingraham/php#Chdir) |
| [scandir](https://www.php.net/manual/en/function.scandir.php) | [Scandir](https://godoc.org/github.com/imartingraham/php#Scandir) |

### Image Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [getimagesize](https://www.php.net/manual/en/function.getimagesize.php) | [GetImageSize](https://godoc.org/github.com/imartingraham/php#GetImageSize) |
| [getimagesizefromstring](https://www.php.net/manual/en/function.getimagesizefromstring.php) | [GetImageSizeFromString](https://godoc.org/github.com/imartingraham/php#GetImageSizeFromString) |

### Network Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [gethostbyaddr](https://php.net/manual/en/function.gethostbyaddr.php) | [GetHostByAddr](https://godoc.org/github.com/imartingraham/php#GetHostByAddr) |
| [gethostbyname](https://php.net/manual/en/function.gethostbyname.php) | [GetHostByName](https://godoc.org/github.com/imartingraham/php#GetHostByName) |
| [gethostbynamel](https://php.net/manual/en/function.gethostbynamel.php) | [GetHostByNamel](https://godoc.org/github.com/imartingraham/php#GetHostByNamel) |
| [gethostname](https://php.net/manual/en/function.gethostname.php) | [GetHostName](https://godoc.org/github.com/imartingraham/php#GetHostName) |
| [ip2long](https://php.net/manual/en/function.ip2long.php) | [IP2Long](https://godoc.org/github.com/imartingraham/php#IP2Long) |
| [long2ip](https://php.net/manual/en/function.long2ip.php) | [Long2IP](https://godoc.org/github.com/imartingraham/php#Long2IP) |

### JSON Functions

| PHP function                             | Golang function                          |
| ---------------------------------------- | ---------------------------------------- |
| [json_decode](https://php.net/manual/en/function.json-decode.php) | [JSONDecode](https://godoc.org/github.com/imartingraham/php#JSONDecode) |
| [json_encode](https://php.net/manual/en/function.json-encode.php) | [JSONEncode](https://godoc.org/github.com/imartingraham/php#JSONEncode) |

### CSPRNG Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [random_bytes](https://php.net/manual/en/function.random-bytes.php) | [RandomBytes](https://godoc.org/github.com/imartingraham/php#RandomBytes) |

### URL Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [base64_decode](https://www.php.net/manual/en/function.base64-decode.php) | [Base64Decode](https://godoc.org/github.com/imartingraham/php#Base64Decode) |
| [base64_encode](https://www.php.net/manual/en/function.base64-encode.php) | [Base64Encode](https://godoc.org/github.com/imartingraham/php#Base64Encode) |
| [get_headers](https://www.php.net/manual/en/function.get-headers.php) | [GetHeaders](https://godoc.org/github.com/imartingraham/php#GetHeaders) |
| [get_meta_tags](https://www.php.net/manual/en/function.get-meta-tags.php) | [GetMetaTags](https://godoc.org/github.com/imartingraham/php#GetMetaTags) |
| [http_build_query](https://www.php.net/manual/en/function.http-build-query.php) | [HTTPBuildQuery](https://godoc.org/github.com/imartingraham/php#HTTPBuildQuery) |
| [parse_url](https://www.php.net/manual/en/function.parse-url.php) | [ParseURL](https://godoc.org/github.com/imartingraham/php#ParseURL) |
| [rawurldecode](https://www.php.net/manual/en/function.rawurldecode.php) | [RawURLDecode](https://godoc.org/github.com/imartingraham/php#RawURLDecode) |
| [rawurlencode](https://www.php.net/manual/en/function.rawurlencode.php) | [RawURLEncode](https://godoc.org/github.com/imartingraham/php#RawURLEncode) |
| [urldecode](https://www.php.net/manual/en/function.urldecode.php) | [URLDecode](https://godoc.org/github.com/imartingraham/php#URLDecode) |
| [urlencode](https://www.php.net/manual/en/function.urlencode.php) | [URLEncode](https://godoc.org/github.com/imartingraham/php#URLEncode) |

### Misc Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [getenv](https://php.net/manual/en/function.getenv.php) | [Getenv](https://godoc.org/github.com/imartingraham/php#Getenv) |
| [putenv](https://php.net/manual/en/function.putenv.php) | [Putenv](https://godoc.org/github.com/imartingraham/php#Putenv) |
| [memory_get_usage](https://php.net/manual/en/function.memory-get-usage.php) | [MemoryGetUsage](https://godoc.org/github.com/imartingraham/php#MemoryGetUsage) |

## Install

```
go get github.com/imartingraham/php
```

## Usage

```
import (
    "fmt"
    "github.com/imartingraham/php"
)

// Date/Time functions

fmt.Println(php.Strtotime("2017-07-14 02:40:00")) // output: 1500000000

fmt.Println(php.Strtotime("2017-07-14T10:40:00+08:00")) // output: 1500000000

fmt.Println(php.Date("Y-m-d H:i:s", 1500000000)) // output: 2017-07-14 02:40:00

fmt.Println(php.Date("c", 1500000000)) // output: 2017-07-14T02:40:00+00:00

// String functions

str := "abcdef"

fmt.Println(php.Substr(str, 1, 0)) // bcdef
fmt.Println(php.Substr(str, 1, 3)) // bcd
fmt.Println(php.Substr(str, 0, 4)) // abcd
fmt.Println(php.Substr(str, -1, 1)) // f
fmt.Println(php.Substr(str, 0, -1)) // abcde

// Math functions

fmt.Println(php.Round(5.055, 2)) // 5.06

// Array functions

arr := []string{"1", "1", "2", "3", "a", "ab", "abc", "abc", "abc", "Abc"}
fmt.Println(php.ArrayUnique(arr).([]string)) // [abc Abc 1 2 3 a ab]
fmt.Println(php.InArray("a", arr)) // true

```

*For more usage you can find it out from test files.*
