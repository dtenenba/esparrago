#define USE_RINTERNALS
#include <R.h>
#include <Rinternals.h>
#include "_cgo_export.h"




SEXP sum_int( SEXP x ){
  // if( TYPEOF(x) != INTSXP ) error("expecting an integer vector") ;
  GoSlice gox = { INTEGER(x), LENGTH(x), LENGTH(x) } ;
  // return ScalarInteger( SumInt(gox) ) ;
  return Rf_ScalarInteger( SumInt(gox) ) ;
}

SEXP sum_double( SEXP x ){
  // if( TYPEOF(x) != REALSXP ) error("expecting a numeric vector") ;
  GoSlice gox = { REAL(x), LENGTH(x), LENGTH(x) } ;
  // return ScalarReal( SumDouble(gox) ) ;
  return Rf_ScalarReal( SumDouble(gox) ) ;
}


SEXP IntegerVectorFromGoSlice( void* data, int n ){
  SEXP res = allocVector( INTSXP, n) ;
  memmove( INTEGER(res), data, sizeof(int)*n ) ;
  return res ;
}

SEXP numbers( SEXP n ){
  //if( TYPEOF(n) != INTSXP || LENGTH(n) != 1 ) error("expecting a single integer") ;

  return Numbers( INTEGER(n)[0] ) ;
}

SEXP godouble(SEXP x){
  return Rf_ScalarInteger( DoubleIt( INTEGER(x)[0] ) ) ;
  // return ScalarInteger( DoubleIt (x) ); // fails
}

SEXP foobar(){
  GoString res = Foobar() ;
  // return ScalarString(mkCharLenCE( res.p, res.n, CE_UTF8 )) ;
  return Rf_ScalarString(mkCharLenCE( res.p, res.n, CE_UTF8 )) ;
}

SEXP nbytes(SEXP x){
  SEXP sx = STRING_ELT(x, 0);
  GoString gos = { (char*) CHAR(sx), SHORT_VEC_LENGTH(sx) };
  // Both of these work here: but ScalarInteger does not work
  // in godouble(). So should we always use Rf_ScalarInteger()?
  return Rf_ScalarInteger( Nbytes(gos) ) ;
  // return ScalarInteger( Nbytes(gos) ) ;
}
