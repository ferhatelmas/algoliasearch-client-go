// Code generated by go generate. DO NOT EDIT.

package search

import (
	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

// QueryParams represents all the parameters that can be passed at query time.
type QueryParams struct {
	AttributesToRetrieve                    *opt.AttributesToRetrieveOption                    `json:"attributesToRetrieve,omitempty"`
	RestrictSearchableAttributes            *opt.RestrictSearchableAttributesOption            `json:"restrictSearchableAttributes,omitempty"`
	Filters                                 *opt.FiltersOption                                 `json:"filters,omitempty"`
	FacetFilters                            *opt.FacetFiltersOption                            `json:"facetFilters,omitempty"`
	OptionalFilters                         *opt.OptionalFiltersOption                         `json:"optionalFilters,omitempty"`
	NumericFilters                          *opt.NumericFiltersOption                          `json:"numericFilters,omitempty"`
	TagFilters                              *opt.TagFiltersOption                              `json:"tagFilters,omitempty"`
	SumOrFiltersScores                      *opt.SumOrFiltersScoresOption                      `json:"sumOrFiltersScores,omitempty"`
	Facets                                  *opt.FacetsOption                                  `json:"facets,omitempty"`
	MaxValuesPerFacet                       *opt.MaxValuesPerFacetOption                       `json:"maxValuesPerFacet,omitempty"`
	FacetingAfterDistinct                   *opt.FacetingAfterDistinctOption                   `json:"facetingAfterDistinct,omitempty"`
	SortFacetValuesBy                       *opt.SortFacetValuesByOption                       `json:"sortFacetValuesBy,omitempty"`
	AttributesToHighlight                   *opt.AttributesToHighlightOption                   `json:"attributesToHighlight,omitempty"`
	AttributesToSnippet                     *opt.AttributesToSnippetOption                     `json:"attributesToSnippet,omitempty"`
	HighlightPreTag                         *opt.HighlightPreTagOption                         `json:"highlightPreTag,omitempty"`
	HighlightPostTag                        *opt.HighlightPostTagOption                        `json:"highlightPostTag,omitempty"`
	SnippetEllipsisText                     *opt.SnippetEllipsisTextOption                     `json:"snippetEllipsisText,omitempty"`
	RestrictHighlightAndSnippetArrays       *opt.RestrictHighlightAndSnippetArraysOption       `json:"restrictHighlightAndSnippetArrays,omitempty"`
	Page                                    *opt.PageOption                                    `json:"page,omitempty"`
	HitsPerPage                             *opt.HitsPerPageOption                             `json:"hitsPerPage,omitempty"`
	Offset                                  *opt.OffsetOption                                  `json:"offset,omitempty"`
	Length                                  *opt.LengthOption                                  `json:"length,omitempty"`
	MinWordSizefor1Typo                     *opt.MinWordSizefor1TypoOption                     `json:"minWordSizefor1Typo,omitempty"`
	MinWordSizefor2Typos                    *opt.MinWordSizefor2TyposOption                    `json:"minWordSizefor2Typos,omitempty"`
	TypoTolerance                           *opt.TypoToleranceOption                           `json:"typoTolerance,omitempty"`
	AllowTyposOnNumericTokens               *opt.AllowTyposOnNumericTokensOption               `json:"allowTyposOnNumericTokens,omitempty"`
	DisableTypoToleranceOnAttributes        *opt.DisableTypoToleranceOnAttributesOption        `json:"disableTypoToleranceOnAttributes,omitempty"`
	AroundLatLng                            *opt.AroundLatLngOption                            `json:"aroundLatLng,omitempty"`
	AroundLatLngViaIP                       *opt.AroundLatLngViaIPOption                       `json:"aroundLatLngViaIP,omitempty"`
	AroundRadius                            *opt.AroundRadiusOption                            `json:"aroundRadius,omitempty"`
	AroundPrecision                         *opt.AroundPrecisionOption                         `json:"aroundPrecision,omitempty"`
	MinimumAroundRadius                     *opt.MinimumAroundRadiusOption                     `json:"minimumAroundRadius,omitempty"`
	InsideBoundingBox                       *opt.InsideBoundingBoxOption                       `json:"insideBoundingBox,omitempty"`
	InsidePolygon                           *opt.InsidePolygonOption                           `json:"insidePolygon,omitempty"`
	IgnorePlurals                           *opt.IgnorePluralsOption                           `json:"ignorePlurals,omitempty"`
	RemoveStopWords                         *opt.RemoveStopWordsOption                         `json:"removeStopWords,omitempty"`
	QueryLanguages                          *opt.QueryLanguagesOption                          `json:"queryLanguages,omitempty"`
	QueryType                               *opt.QueryTypeOption                               `json:"queryType,omitempty"`
	RemoveWordsIfNoResults                  *opt.RemoveWordsIfNoResultsOption                  `json:"removeWordsIfNoResults,omitempty"`
	AdvancedSyntax                          *opt.AdvancedSyntaxOption                          `json:"advancedSyntax,omitempty"`
	OptionalWords                           *opt.OptionalWordsOption                           `json:"optionalWords,omitempty"`
	DisableExactOnAttributes                *opt.DisableExactOnAttributesOption                `json:"disableExactOnAttributes,omitempty"`
	ExactOnSingleWordQuery                  *opt.ExactOnSingleWordQueryOption                  `json:"exactOnSingleWordQuery,omitempty"`
	AlternativesAsExact                     *opt.AlternativesAsExactOption                     `json:"alternativesAsExact,omitempty"`
	AdvancedSyntaxFeatures                  *opt.AdvancedSyntaxFeaturesOption                  `json:"advancedSyntaxFeatures,omitempty"`
	SimilarQuery                            *opt.SimilarQueryOption                            `json:"similarQuery,omitempty"`
	EnableABTest                            *opt.EnableABTestOption                            `json:"enableABTest,omitempty"`
	EnableRules                             *opt.EnableRulesOption                             `json:"enableRules,omitempty"`
	RuleContexts                            *opt.RuleContextsOption                            `json:"ruleContexts,omitempty"`
	EnablePersonalization                   *opt.EnablePersonalizationOption                   `json:"enablePersonalization,omitempty"`
	PersonalizationImpact                   *opt.PersonalizationImpactOption                   `json:"personalizationImpact,omitempty"`
	UserToken                               *opt.UserTokenOption                               `json:"userToken,omitempty"`
	Distinct                                *opt.DistinctOption                                `json:"distinct,omitempty"`
	GetRankingInfo                          *opt.GetRankingInfoOption                          `json:"getRankingInfo,omitempty"`
	ClickAnalytics                          *opt.ClickAnalyticsOption                          `json:"clickAnalytics,omitempty"`
	Analytics                               *opt.AnalyticsOption                               `json:"analytics,omitempty"`
	AnalyticsTags                           *opt.AnalyticsTagsOption                           `json:"analyticsTags,omitempty"`
	Synonyms                                *opt.SynonymsOption                                `json:"synonyms,omitempty"`
	ReplaceSynonymsInHighlight              *opt.ReplaceSynonymsInHighlightOption              `json:"replaceSynonymsInHighlight,omitempty"`
	MinProximity                            *opt.MinProximityOption                            `json:"minProximity,omitempty"`
	ResponseFields                          *opt.ResponseFieldsOption                          `json:"responseFields,omitempty"`
	MaxFacetHits                            *opt.MaxFacetHitsOption                            `json:"maxFacetHits,omitempty"`
	PercentileComputation                   *opt.PercentileComputationOption                   `json:"percentileComputation,omitempty"`
	Explain                                 *opt.ExplainOption                                 `json:"explain,omitempty"`
	AttributeCriteriaComputedByMinProximity *opt.AttributeCriteriaComputedByMinProximityOption `json:"attributeCriteriaComputedByMinProximity,omitempty"`
}

func newQueryParams(opts ...interface{}) QueryParams {
	return QueryParams{
		AttributesToRetrieve:                    iopt.ExtractAttributesToRetrieve(opts...),
		RestrictSearchableAttributes:            iopt.ExtractRestrictSearchableAttributes(opts...),
		Filters:                                 iopt.ExtractFilters(opts...),
		FacetFilters:                            iopt.ExtractFacetFilters(opts...),
		OptionalFilters:                         iopt.ExtractOptionalFilters(opts...),
		NumericFilters:                          iopt.ExtractNumericFilters(opts...),
		TagFilters:                              iopt.ExtractTagFilters(opts...),
		SumOrFiltersScores:                      iopt.ExtractSumOrFiltersScores(opts...),
		Facets:                                  iopt.ExtractFacets(opts...),
		MaxValuesPerFacet:                       iopt.ExtractMaxValuesPerFacet(opts...),
		FacetingAfterDistinct:                   iopt.ExtractFacetingAfterDistinct(opts...),
		SortFacetValuesBy:                       iopt.ExtractSortFacetValuesBy(opts...),
		AttributesToHighlight:                   iopt.ExtractAttributesToHighlight(opts...),
		AttributesToSnippet:                     iopt.ExtractAttributesToSnippet(opts...),
		HighlightPreTag:                         iopt.ExtractHighlightPreTag(opts...),
		HighlightPostTag:                        iopt.ExtractHighlightPostTag(opts...),
		SnippetEllipsisText:                     iopt.ExtractSnippetEllipsisText(opts...),
		RestrictHighlightAndSnippetArrays:       iopt.ExtractRestrictHighlightAndSnippetArrays(opts...),
		Page:                                    iopt.ExtractPage(opts...),
		HitsPerPage:                             iopt.ExtractHitsPerPage(opts...),
		Offset:                                  iopt.ExtractOffset(opts...),
		Length:                                  iopt.ExtractLength(opts...),
		MinWordSizefor1Typo:                     iopt.ExtractMinWordSizefor1Typo(opts...),
		MinWordSizefor2Typos:                    iopt.ExtractMinWordSizefor2Typos(opts...),
		TypoTolerance:                           iopt.ExtractTypoTolerance(opts...),
		AllowTyposOnNumericTokens:               iopt.ExtractAllowTyposOnNumericTokens(opts...),
		DisableTypoToleranceOnAttributes:        iopt.ExtractDisableTypoToleranceOnAttributes(opts...),
		AroundLatLng:                            iopt.ExtractAroundLatLng(opts...),
		AroundLatLngViaIP:                       iopt.ExtractAroundLatLngViaIP(opts...),
		AroundRadius:                            iopt.ExtractAroundRadius(opts...),
		AroundPrecision:                         iopt.ExtractAroundPrecision(opts...),
		MinimumAroundRadius:                     iopt.ExtractMinimumAroundRadius(opts...),
		InsideBoundingBox:                       iopt.ExtractInsideBoundingBox(opts...),
		InsidePolygon:                           iopt.ExtractInsidePolygon(opts...),
		IgnorePlurals:                           iopt.ExtractIgnorePlurals(opts...),
		RemoveStopWords:                         iopt.ExtractRemoveStopWords(opts...),
		QueryLanguages:                          iopt.ExtractQueryLanguages(opts...),
		QueryType:                               iopt.ExtractQueryType(opts...),
		RemoveWordsIfNoResults:                  iopt.ExtractRemoveWordsIfNoResults(opts...),
		AdvancedSyntax:                          iopt.ExtractAdvancedSyntax(opts...),
		OptionalWords:                           iopt.ExtractOptionalWords(opts...),
		DisableExactOnAttributes:                iopt.ExtractDisableExactOnAttributes(opts...),
		ExactOnSingleWordQuery:                  iopt.ExtractExactOnSingleWordQuery(opts...),
		AlternativesAsExact:                     iopt.ExtractAlternativesAsExact(opts...),
		AdvancedSyntaxFeatures:                  iopt.ExtractAdvancedSyntaxFeatures(opts...),
		SimilarQuery:                            iopt.ExtractSimilarQuery(opts...),
		EnableABTest:                            iopt.ExtractEnableABTest(opts...),
		EnableRules:                             iopt.ExtractEnableRules(opts...),
		RuleContexts:                            iopt.ExtractRuleContexts(opts...),
		EnablePersonalization:                   iopt.ExtractEnablePersonalization(opts...),
		PersonalizationImpact:                   iopt.ExtractPersonalizationImpact(opts...),
		UserToken:                               iopt.ExtractUserToken(opts...),
		Distinct:                                iopt.ExtractDistinct(opts...),
		GetRankingInfo:                          iopt.ExtractGetRankingInfo(opts...),
		ClickAnalytics:                          iopt.ExtractClickAnalytics(opts...),
		Analytics:                               iopt.ExtractAnalytics(opts...),
		AnalyticsTags:                           iopt.ExtractAnalyticsTags(opts...),
		Synonyms:                                iopt.ExtractSynonyms(opts...),
		ReplaceSynonymsInHighlight:              iopt.ExtractReplaceSynonymsInHighlight(opts...),
		MinProximity:                            iopt.ExtractMinProximity(opts...),
		ResponseFields:                          iopt.ExtractResponseFields(opts...),
		MaxFacetHits:                            iopt.ExtractMaxFacetHits(opts...),
		PercentileComputation:                   iopt.ExtractPercentileComputation(opts...),
		Explain:                                 iopt.ExtractExplain(opts...),
		AttributeCriteriaComputedByMinProximity: iopt.ExtractAttributeCriteriaComputedByMinProximity(opts...),
	}
}

func (p *QueryParams) Equal(p2 *QueryParams) bool {
	if p == nil && p2 == nil {
		return true
	}

	if p != nil && p2 == nil || p == nil && p2 != nil {
		return false
	}

	if !p.AttributesToRetrieve.Equal(p2.AttributesToRetrieve) {
		return false
	}
	if !p.RestrictSearchableAttributes.Equal(p2.RestrictSearchableAttributes) {
		return false
	}
	if !p.Filters.Equal(p2.Filters) {
		return false
	}
	if !p.FacetFilters.Equal(p2.FacetFilters) {
		return false
	}
	if !p.OptionalFilters.Equal(p2.OptionalFilters) {
		return false
	}
	if !p.NumericFilters.Equal(p2.NumericFilters) {
		return false
	}
	if !p.TagFilters.Equal(p2.TagFilters) {
		return false
	}
	if !p.SumOrFiltersScores.Equal(p2.SumOrFiltersScores) {
		return false
	}
	if !p.Facets.Equal(p2.Facets) {
		return false
	}
	if !p.MaxValuesPerFacet.Equal(p2.MaxValuesPerFacet) {
		return false
	}
	if !p.FacetingAfterDistinct.Equal(p2.FacetingAfterDistinct) {
		return false
	}
	if !p.SortFacetValuesBy.Equal(p2.SortFacetValuesBy) {
		return false
	}
	if !p.AttributesToHighlight.Equal(p2.AttributesToHighlight) {
		return false
	}
	if !p.AttributesToSnippet.Equal(p2.AttributesToSnippet) {
		return false
	}
	if !p.HighlightPreTag.Equal(p2.HighlightPreTag) {
		return false
	}
	if !p.HighlightPostTag.Equal(p2.HighlightPostTag) {
		return false
	}
	if !p.SnippetEllipsisText.Equal(p2.SnippetEllipsisText) {
		return false
	}
	if !p.RestrictHighlightAndSnippetArrays.Equal(p2.RestrictHighlightAndSnippetArrays) {
		return false
	}
	if !p.Page.Equal(p2.Page) {
		return false
	}
	if !p.HitsPerPage.Equal(p2.HitsPerPage) {
		return false
	}
	if !p.Offset.Equal(p2.Offset) {
		return false
	}
	if !p.Length.Equal(p2.Length) {
		return false
	}
	if !p.MinWordSizefor1Typo.Equal(p2.MinWordSizefor1Typo) {
		return false
	}
	if !p.MinWordSizefor2Typos.Equal(p2.MinWordSizefor2Typos) {
		return false
	}
	if !p.TypoTolerance.Equal(p2.TypoTolerance) {
		return false
	}
	if !p.AllowTyposOnNumericTokens.Equal(p2.AllowTyposOnNumericTokens) {
		return false
	}
	if !p.DisableTypoToleranceOnAttributes.Equal(p2.DisableTypoToleranceOnAttributes) {
		return false
	}
	if !p.AroundLatLng.Equal(p2.AroundLatLng) {
		return false
	}
	if !p.AroundLatLngViaIP.Equal(p2.AroundLatLngViaIP) {
		return false
	}
	if !p.AroundRadius.Equal(p2.AroundRadius) {
		return false
	}
	if !p.AroundPrecision.Equal(p2.AroundPrecision) {
		return false
	}
	if !p.MinimumAroundRadius.Equal(p2.MinimumAroundRadius) {
		return false
	}
	if !p.InsideBoundingBox.Equal(p2.InsideBoundingBox) {
		return false
	}
	if !p.InsidePolygon.Equal(p2.InsidePolygon) {
		return false
	}
	if !p.IgnorePlurals.Equal(p2.IgnorePlurals) {
		return false
	}
	if !p.RemoveStopWords.Equal(p2.RemoveStopWords) {
		return false
	}
	if !p.QueryLanguages.Equal(p2.QueryLanguages) {
		return false
	}
	if !p.QueryType.Equal(p2.QueryType) {
		return false
	}
	if !p.RemoveWordsIfNoResults.Equal(p2.RemoveWordsIfNoResults) {
		return false
	}
	if !p.AdvancedSyntax.Equal(p2.AdvancedSyntax) {
		return false
	}
	if !p.OptionalWords.Equal(p2.OptionalWords) {
		return false
	}
	if !p.DisableExactOnAttributes.Equal(p2.DisableExactOnAttributes) {
		return false
	}
	if !p.ExactOnSingleWordQuery.Equal(p2.ExactOnSingleWordQuery) {
		return false
	}
	if !p.AlternativesAsExact.Equal(p2.AlternativesAsExact) {
		return false
	}
	if !p.AdvancedSyntaxFeatures.Equal(p2.AdvancedSyntaxFeatures) {
		return false
	}
	if !p.SimilarQuery.Equal(p2.SimilarQuery) {
		return false
	}
	if !p.EnableABTest.Equal(p2.EnableABTest) {
		return false
	}
	if !p.EnableRules.Equal(p2.EnableRules) {
		return false
	}
	if !p.RuleContexts.Equal(p2.RuleContexts) {
		return false
	}
	if !p.EnablePersonalization.Equal(p2.EnablePersonalization) {
		return false
	}
	if !p.PersonalizationImpact.Equal(p2.PersonalizationImpact) {
		return false
	}
	if !p.UserToken.Equal(p2.UserToken) {
		return false
	}
	if !p.Distinct.Equal(p2.Distinct) {
		return false
	}
	if !p.GetRankingInfo.Equal(p2.GetRankingInfo) {
		return false
	}
	if !p.ClickAnalytics.Equal(p2.ClickAnalytics) {
		return false
	}
	if !p.Analytics.Equal(p2.Analytics) {
		return false
	}
	if !p.AnalyticsTags.Equal(p2.AnalyticsTags) {
		return false
	}
	if !p.Synonyms.Equal(p2.Synonyms) {
		return false
	}
	if !p.ReplaceSynonymsInHighlight.Equal(p2.ReplaceSynonymsInHighlight) {
		return false
	}
	if !p.MinProximity.Equal(p2.MinProximity) {
		return false
	}
	if !p.ResponseFields.Equal(p2.ResponseFields) {
		return false
	}
	if !p.MaxFacetHits.Equal(p2.MaxFacetHits) {
		return false
	}
	if !p.PercentileComputation.Equal(p2.PercentileComputation) {
		return false
	}
	if !p.Explain.Equal(p2.Explain) {
		return false
	}
	if !p.AttributeCriteriaComputedByMinProximity.Equal(p2.AttributeCriteriaComputedByMinProximity) {
		return false
	}

	return true
}
