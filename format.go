// Format
package util

import (
	"math"
	"reflect"
	"runtime"
	"strconv"
	//	"fmt"
)

const CalculationError float64 = 1e-9

func FormatOutput(o interface{}) {

	//	t := reflect.TypeOf(o)
	//	if k := t.Kind(); k != reflect.Struct {
	//		return
	//	}
	//	var j int
	v := reflect.ValueOf(o)
	if v.Kind() != reflect.Ptr || !v.Elem().CanSet() {
		return
	} else {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {

		case reflect.Float32:
			if (math.Abs(v.Field(i).Float()) - 1) > CalculationError {
				a := strconv.FormatFloat(v.Field(i).Float(), 'f', 2, 32)
				f, _ := strconv.ParseFloat(a, 32)
				v.Field(i).SetFloat(f)
			} else {
				a := strconv.FormatFloat(v.Field(i).Float(), 'e', 2, 32)
				f, _ := strconv.ParseFloat(a, 32)
				v.Field(i).SetFloat(f)
			}
		case reflect.Float64:
			if (math.Abs(v.Field(i).Float()) - 1) > CalculationError {
				a := strconv.FormatFloat(v.Field(i).Float(), 'f', 2, 64)
				f, _ := strconv.ParseFloat(a, 64)
				v.Field(i).SetFloat(f)
			} else {
				a := strconv.FormatFloat(v.Field(i).Float(), 'e', 2, 64)
				f, _ := strconv.ParseFloat(a, 64)
				v.Field(i).SetFloat(f)
			}

		case reflect.Struct:
			for j := 0; j < v.Field(i).NumField(); j++ {
				switch v.Field(i).Field(j).Kind() {
				case reflect.Float32:
					if (math.Abs(v.Field(i).Field(j).Float()) - 1) > CalculationError {
						a := strconv.FormatFloat(v.Field(i).Field(j).Float(), 'f', 2, 32)
						f, _ := strconv.ParseFloat(a, 32)
						v.Field(i).Field(j).SetFloat(f)
					} else {
						a := strconv.FormatFloat(v.Field(i).Field(j).Float(), 'e', 2, 32)
						f, _ := strconv.ParseFloat(a, 32)
						v.Field(i).Field(j).SetFloat(f)
					}
				case reflect.Float64:
					if (math.Abs(v.Field(i).Field(j).Float()) - 1) > CalculationError {
						a := strconv.FormatFloat(v.Field(i).Field(j).Float(), 'f', 2, 64)
						f, _ := strconv.ParseFloat(a, 64)
						v.Field(i).Field(j).SetFloat(f)
					} else {
						a := strconv.FormatFloat(v.Field(i).Field(j).Float(), 'e', 2, 64)
						f, _ := strconv.ParseFloat(a, 64)
						v.Field(i).Field(j).SetFloat(f)
					}
				case reflect.Struct:
					for k := 0; k < v.Field(i).Field(j).NumField(); k++ {
						switch v.Field(i).Field(j).Field(k).Kind() {
						case reflect.Float32:
							if (math.Abs(v.Field(i).Field(j).Field(k).Float()) - 1) > CalculationError {
								a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Float(), 'f', 2, 32)
								f, _ := strconv.ParseFloat(a, 32)
								v.Field(i).Field(j).Field(k).SetFloat(f)
							} else {
								a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Float(), 'e', 2, 32)
								f, _ := strconv.ParseFloat(a, 32)
								v.Field(i).Field(j).Field(k).SetFloat(f)
							}
						case reflect.Float64:
							if (math.Abs(v.Field(i).Field(j).Field(k).Float()) - 1) > CalculationError {
								a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Float(), 'f', 2, 64)
								f, _ := strconv.ParseFloat(a, 64)
								v.Field(i).Field(j).Field(k).SetFloat(f)
							} else {
								a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Float(), 'e', 2, 64)
								f, _ := strconv.ParseFloat(a, 64)
								v.Field(i).Field(j).Field(k).SetFloat(f)
							}
						case reflect.Slice:
							for m := 0; m < v.Field(i).Field(j).Field(k).Len(); m++ {
								switch v.Field(i).Field(j).Field(k).Index(m).Kind() {
								case reflect.Float32:
									if (math.Abs(v.Field(i).Field(j).Field(k).Index(m).Float()) - 1) > CalculationError {
										a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Index(m).Float(), 'f', 2, 32)
										f, _ := strconv.ParseFloat(a, 32)
										v.Field(i).Field(j).Field(k).Index(m).SetFloat(f)
									} else {
										a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Index(m).Float(), 'e', 2, 32)
										f, _ := strconv.ParseFloat(a, 32)
										v.Field(i).Field(j).Field(k).Index(m).SetFloat(f)
									}
								case reflect.Float64:
									if (math.Abs(v.Field(i).Field(j).Field(k).Index(m).Float()) - 1) > CalculationError {
										a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Index(m).Float(), 'f', 2, 64)
										f, _ := strconv.ParseFloat(a, 64)
										v.Field(i).Field(j).Field(k).Index(m).SetFloat(f)
									} else {
										a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Index(m).Float(), 'e', 2, 64)
										f, _ := strconv.ParseFloat(a, 64)
										v.Field(i).Field(j).Field(k).Index(m).SetFloat(f)
									}
								case reflect.Struct:
									for n := 0; n < v.Field(i).Field(j).Field(k).Index(m).NumField(); n++ {
										switch v.Field(i).Field(j).Field(k).Index(m).Field(n).Kind() {
										case reflect.Float32:
											if (math.Abs(v.Field(i).Field(j).Field(k).Index(m).Field(n).Float()) - 1) > CalculationError {
												a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Index(m).Field(n).Float(), 'f', 2, 32)
												f, _ := strconv.ParseFloat(a, 32)
												v.Field(i).Field(j).Field(k).Index(m).Field(n).SetFloat(f)
											} else {
												a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Index(m).Field(n).Float(), 'e', 2, 32)
												f, _ := strconv.ParseFloat(a, 32)
												v.Field(i).Field(j).Field(k).Index(m).Field(n).SetFloat(f)
											}
										case reflect.Float64:
											if (math.Abs(v.Field(i).Field(j).Field(k).Index(m).Field(n).Float()) - 1) > CalculationError {
												a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Index(m).Field(n).Float(), 'f', 2, 64)
												f, _ := strconv.ParseFloat(a, 64)
												v.Field(i).Field(j).Field(k).Index(m).Field(n).SetFloat(f)
											} else {
												a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Index(m).Field(n).Float(), 'e', 2, 64)
												f, _ := strconv.ParseFloat(a, 64)
												v.Field(i).Field(j).Field(k).Index(m).Field(n).SetFloat(f)
											}
										}
										runtime.Gosched()
									}

								}
								runtime.Gosched()
							}

						case reflect.Struct:
							for m := 0; m < v.Field(i).Field(j).Field(k).NumField(); m++ {
								switch v.Field(i).Field(j).Field(k).Field(m).Kind() {
								case reflect.Float32:
									if (math.Abs(v.Field(i).Field(j).Field(k).Field(m).Float()) - 1) > CalculationError {
										a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Field(m).Float(), 'f', 2, 32)
										f, _ := strconv.ParseFloat(a, 32)
										v.Field(i).Field(j).Field(k).Field(m).SetFloat(f)
									} else {
										a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Field(m).Float(), 'e', 2, 32)
										f, _ := strconv.ParseFloat(a, 32)
										v.Field(i).Field(j).Field(k).Field(m).SetFloat(f)
									}
								case reflect.Float64:
									if (math.Abs(v.Field(i).Field(j).Field(k).Field(m).Float()) - 1) > CalculationError {
										a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Field(m).Float(), 'f', 2, 64)
										f, _ := strconv.ParseFloat(a, 64)
										v.Field(i).Field(j).Field(k).Field(m).SetFloat(f)
									} else {
										a := strconv.FormatFloat(v.Field(i).Field(j).Field(k).Field(m).Float(), 'e', 2, 64)
										f, _ := strconv.ParseFloat(a, 64)
										v.Field(i).Field(j).Field(k).Field(m).SetFloat(f)
									}
								}
								runtime.Gosched()
							}

						}
						runtime.Gosched()
					}
				case reflect.Slice:
					for k := 0; k < v.Field(i).Field(j).Len(); k++ {
						switch v.Field(i).Field(j).Index(k).Kind() {

						case reflect.Float32:
							if (math.Abs(v.Field(i).Field(j).Index(k).Float()) - 1) > CalculationError {
								a := strconv.FormatFloat(v.Field(i).Field(j).Index(k).Float(), 'f', 2, 32)
								f, _ := strconv.ParseFloat(a, 32)
								v.Field(i).Field(j).Index(k).SetFloat(f)
							} else {
								a := strconv.FormatFloat(v.Field(i).Field(j).Index(k).Float(), 'e', 2, 32)
								f, _ := strconv.ParseFloat(a, 32)
								v.Field(i).Field(j).Index(k).SetFloat(f)
							}
						case reflect.Float64:
							if (math.Abs(v.Field(i).Field(j).Index(k).Float()) - 1) > CalculationError {
								a := strconv.FormatFloat(v.Field(i).Field(j).Index(k).Float(), 'f', 2, 64)
								f, _ := strconv.ParseFloat(a, 64)
								v.Field(i).Field(j).Index(k).SetFloat(f)
							} else {
								a := strconv.FormatFloat(v.Field(i).Field(j).Index(k).Float(), 'e', 2, 64)
								f, _ := strconv.ParseFloat(a, 64)
								v.Field(i).Field(j).Index(k).SetFloat(f)
							}
						case reflect.Struct:
							for m := 0; m < v.Field(i).Field(j).Index(k).NumField(); m++ {
								switch v.Field(i).Field(j).Index(k).Field(m).Kind() {
								case reflect.Float32:
									if (math.Abs(v.Field(i).Field(j).Index(k).Field(m).Float()) - 1) > CalculationError {
										a := strconv.FormatFloat(v.Field(i).Field(j).Index(k).Field(m).Float(), 'f', 2, 32)
										f, _ := strconv.ParseFloat(a, 32)
										v.Field(i).Field(j).Index(k).Field(m).SetFloat(f)
									} else {
										a := strconv.FormatFloat(v.Field(i).Field(j).Index(k).Field(m).Float(), 'e', 2, 32)
										f, _ := strconv.ParseFloat(a, 32)
										v.Field(i).Field(j).Index(k).Field(m).SetFloat(f)
									}
								case reflect.Float64:
									if (math.Abs(v.Field(i).Field(j).Index(k).Field(m).Float()) - 1) > CalculationError {
										a := strconv.FormatFloat(v.Field(i).Field(j).Index(k).Field(m).Float(), 'f', 2, 64)
										f, _ := strconv.ParseFloat(a, 64)
										v.Field(i).Field(j).Index(k).Field(m).SetFloat(f)
									} else {
										a := strconv.FormatFloat(v.Field(i).Field(j).Index(k).Field(m).Float(), 'e', 2, 64)
										f, _ := strconv.ParseFloat(a, 64)
										v.Field(i).Field(j).Index(k).Field(m).SetFloat(f)
									}
								}
								runtime.Gosched()
							}
						}
						runtime.Gosched()
					}
				}
				runtime.Gosched()
			}

		case reflect.Slice:
			for j := 0; j < v.Field(i).Len(); j++ {
				switch v.Field(i).Index(j).Kind() {
				case reflect.Float32:
					if (math.Abs(v.Field(i).Index(j).Float()) - 1) > CalculationError {
						a := strconv.FormatFloat(v.Field(i).Index(j).Float(), 'f', 2, 32)
						f, _ := strconv.ParseFloat(a, 32)
						v.Field(i).Index(j).SetFloat(f)
					} else {
						a := strconv.FormatFloat(v.Field(i).Index(j).Float(), 'e', 2, 32)
						f, _ := strconv.ParseFloat(a, 32)
						v.Field(i).Index(j).SetFloat(f)
					}
				case reflect.Float64:
					if (math.Abs(v.Field(i).Index(j).Float()) - 1) > CalculationError {
						a := strconv.FormatFloat(v.Field(i).Index(j).Float(), 'f', 2, 64)
						f, _ := strconv.ParseFloat(a, 64)
						v.Field(i).Index(j).SetFloat(f)
					} else {
						a := strconv.FormatFloat(v.Field(i).Index(j).Float(), 'e', 2, 64)
						f, _ := strconv.ParseFloat(a, 64)
						v.Field(i).Index(j).SetFloat(f)
					}
				case reflect.Slice:
					for k := 0; k < v.Field(i).Index(j).Len(); k++ {
						switch v.Field(i).Index(j).Index(k).Kind() {
						case reflect.Float32:
							if (math.Abs(v.Field(i).Index(j).Index(k).Float()) - 1) > CalculationError {
								a := strconv.FormatFloat(v.Field(i).Index(j).Index(k).Float(), 'f', 2, 32)
								f, _ := strconv.ParseFloat(a, 32)
								v.Field(i).Index(j).Index(k).SetFloat(f)
							} else {
								a := strconv.FormatFloat(v.Field(i).Index(j).Index(k).Float(), 'e', 2, 32)
								f, _ := strconv.ParseFloat(a, 32)
								v.Field(i).Index(j).Index(k).SetFloat(f)
							}
						case reflect.Float64:
							if (math.Abs(v.Field(i).Index(j).Index(k).Float()) - 1) > CalculationError {
								a := strconv.FormatFloat(v.Field(i).Index(j).Index(k).Float(), 'f', 2, 64)
								f, _ := strconv.ParseFloat(a, 64)
								v.Field(i).Index(j).Index(k).SetFloat(f)
							} else {
								a := strconv.FormatFloat(v.Field(i).Index(j).Index(k).Float(), 'e', 2, 64)
								f, _ := strconv.ParseFloat(a, 64)
								v.Field(i).Index(j).Index(k).SetFloat(f)
							}
						}
						runtime.Gosched()
					}
				case reflect.Array:
					for k := 0; k < v.Field(i).Index(j).Len(); k++ {
						switch v.Field(i).Index(j).Index(k).Kind() {
						case reflect.Float32:
							if (math.Abs(v.Field(i).Index(j).Index(k).Float()) - 1) > CalculationError {
								a := strconv.FormatFloat(v.Field(i).Index(j).Index(k).Float(), 'f', 2, 32)
								f, _ := strconv.ParseFloat(a, 32)
								v.Field(i).Index(j).Index(k).SetFloat(f)
							} else {
								a := strconv.FormatFloat(v.Field(i).Index(j).Index(k).Float(), 'e', 2, 32)
								f, _ := strconv.ParseFloat(a, 32)
								v.Field(i).Index(j).Index(k).SetFloat(f)
							}
						case reflect.Float64:
							if (math.Abs(v.Field(i).Index(j).Index(k).Float()) - 1) > CalculationError {
								a := strconv.FormatFloat(v.Field(i).Index(j).Index(k).Float(), 'f', 2, 64)
								f, _ := strconv.ParseFloat(a, 64)
								v.Field(i).Index(j).Index(k).SetFloat(f)
							} else {
								a := strconv.FormatFloat(v.Field(i).Index(j).Index(k).Float(), 'e', 2, 64)
								f, _ := strconv.ParseFloat(a, 64)
								v.Field(i).Index(j).Index(k).SetFloat(f)
							}
						}
						runtime.Gosched()
					}
				case reflect.Struct:
					for k := 0; k < v.Field(i).Index(j).NumField(); k++ {
						switch v.Field(i).Index(j).Field(k).Kind() {
						case reflect.Float32:
							if (math.Abs(v.Field(i).Index(j).Field(k).Float()) - 1) > CalculationError {
								a := strconv.FormatFloat(v.Field(i).Index(j).Field(k).Float(), 'f', 2, 32)
								f, _ := strconv.ParseFloat(a, 32)
								v.Field(i).Index(j).Field(k).SetFloat(f)
							} else {
								a := strconv.FormatFloat(v.Field(i).Index(j).Field(k).Float(), 'e', 2, 32)
								f, _ := strconv.ParseFloat(a, 32)
								v.Field(i).Index(j).Field(k).SetFloat(f)
							}
						case reflect.Float64:
							if (math.Abs(v.Field(i).Index(j).Field(k).Float()) - 1) > CalculationError {
								a := strconv.FormatFloat(v.Field(i).Index(j).Field(k).Float(), 'f', 2, 64)
								f, _ := strconv.ParseFloat(a, 64)
								v.Field(i).Index(j).Field(k).SetFloat(f)
							} else {
								a := strconv.FormatFloat(v.Field(i).Index(j).Field(k).Float(), 'e', 2, 64)
								f, _ := strconv.ParseFloat(a, 64)
								v.Field(i).Index(j).Field(k).SetFloat(f)
							}
						case reflect.Struct:
							for l := 0; l < v.Field(i).Index(j).Field(k).NumField(); l++ {
								switch v.Field(i).Index(j).Field(k).Field(l).Kind() {
								case reflect.Float32:
									if (math.Abs(v.Field(i).Index(j).Field(k).Field(l).Float()) - 1) > CalculationError {
										a := strconv.FormatFloat(v.Field(i).Index(j).Field(k).Field(l).Float(), 'f', 2, 32)
										f, _ := strconv.ParseFloat(a, 32)
										v.Field(i).Index(j).Field(k).Field(l).SetFloat(f)
									} else {
										a := strconv.FormatFloat(v.Field(i).Index(j).Field(k).Field(l).Float(), 'e', 2, 32)
										f, _ := strconv.ParseFloat(a, 32)
										v.Field(i).Index(j).Field(k).Field(l).SetFloat(f)
									}
								case reflect.Float64:
									if (math.Abs(v.Field(i).Index(j).Field(k).Field(l).Float()) - 1) > CalculationError {
										a := strconv.FormatFloat(v.Field(i).Index(j).Field(k).Field(l).Float(), 'f', 2, 64)
										f, _ := strconv.ParseFloat(a, 64)
										v.Field(i).Index(j).Field(k).Field(l).SetFloat(f)
									} else {
										a := strconv.FormatFloat(v.Field(i).Index(j).Field(k).Field(l).Float(), 'e', 2, 64)
										f, _ := strconv.ParseFloat(a, 64)
										v.Field(i).Index(j).Field(k).Field(l).SetFloat(f)
									}
								}
								runtime.Gosched()
							}

						}
						runtime.Gosched()
					}
				}
				runtime.Gosched()
			}

		case reflect.Array:
			for j := 0; j < v.Field(i).Len(); j++ {
				switch v.Field(i).Index(j).Kind() {
				case reflect.Float32:
					if (math.Abs(v.Field(i).Index(j).Float()) - 1) > CalculationError {
						a := strconv.FormatFloat(v.Field(i).Index(j).Float(), 'f', 2, 32)
						f, _ := strconv.ParseFloat(a, 32)
						v.Field(i).Index(j).SetFloat(f)
					} else {
						a := strconv.FormatFloat(v.Field(i).Index(j).Float(), 'e', 2, 32)
						f, _ := strconv.ParseFloat(a, 32)
						v.Field(i).Index(j).SetFloat(f)
					}
				case reflect.Float64:
					if (math.Abs(v.Field(i).Index(j).Float()) - 1) > CalculationError {
						a := strconv.FormatFloat(v.Field(i).Index(j).Float(), 'f', 2, 64)
						f, _ := strconv.ParseFloat(a, 64)
						v.Field(i).Index(j).SetFloat(f)
					} else {
						a := strconv.FormatFloat(v.Field(i).Index(j).Float(), 'e', 2, 64)
						f, _ := strconv.ParseFloat(a, 64)
						v.Field(i).Index(j).SetFloat(f)
					}

					//				case reflect.Slice:
					//					for k := 0; k < v.Field(i).Index(j).Len(); k++ {
					//						switch v.Field(i).Index(j).Index(k).Kind() {
					//						case reflect.Float64:
					//							if (math.Abs(v.Field(i).Index(j).Index(k).Float()) - 1) > CalculationError {
					//								a := strconv.FormatFloat(v.Field(i).Index(j).Index(k).Float(), 'f', 2, 64)
					//								f, _ := strconv.ParseFloat(a, 64)
					//								v.Field(i).Index(j).Index(k).SetFloat(f)
					//							} else {
					//								a := strconv.FormatFloat(v.Field(i).Index(j).Index(k).Float(), 'e', 2, 64)
					//								f, _ := strconv.ParseFloat(a, 64)
					//								v.Field(i).Index(j).Index(k).SetFloat(f)
					//							}
					//						}
					//					}
					//				case reflect.Array:
					//					for k := 0; k < v.Field(i).Index(j).Len(); k++ {
					//						switch v.Field(i).Index(j).Index(k).Kind() {
					//						case reflect.Float64:
					//							if (math.Abs(v.Field(i).Index(j).Index(k).Float()) - 1) > CalculationError {
					//								a := strconv.FormatFloat(v.Field(i).Index(j).Index(k).Float(), 'f', 2, 64)
					//								f, _ := strconv.ParseFloat(a, 64)
					//								v.Field(i).Index(j).Index(k).SetFloat(f)
					//							} else {
					//								a := strconv.FormatFloat(v.Field(i).Index(j).Index(k).Float(), 'e', 2, 64)
					//								f, _ := strconv.ParseFloat(a, 64)
					//								v.Field(i).Index(j).Index(k).SetFloat(f)
					//							}
					//						}
					//					}
				}
				runtime.Gosched()
			}

		}
		runtime.Gosched()
	}

}

//函数将浮点数表示为字符串并返回。

//bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。

//fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。

//prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
