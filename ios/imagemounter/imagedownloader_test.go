package imagemounter_test

import (
	"github.com/danielpaulus/go-ios/ios/imagemounter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionMatching(t *testing.T) {
	assert.Equal(t, "11.2", imagemounter.MatchAvailable("11.2.5"))
	assert.Equal(t, "12.2 (16E226)", imagemounter.MatchAvailable("12.2.5"))
	assert.Equal(t, "13.6", imagemounter.MatchAvailable("13.6.1"))
	assert.Equal(t, "14.7", imagemounter.MatchAvailable("14.7.1"))
	assert.Equal(t, "15.2", imagemounter.MatchAvailable("15.3.1"))
	assert.Equal(t, "15.4", imagemounter.MatchAvailable("15.4.1"))
	assert.Equal(t, "15.4", imagemounter.MatchAvailable("19.4.1"))
}
