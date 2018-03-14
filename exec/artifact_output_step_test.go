package exec_test

import (
	"bytes"
	"context"

	"github.com/concourse/atc/exec"
	"github.com/concourse/atc/worker/workerfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ArtifactOutputStep", func() {
	var (
		ctx    context.Context
		cancel func()

		state exec.RunState

		step    exec.Step
		stepErr error
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())

		state = exec.NewRunState()
	})

	JustBeforeEach(func() {
		step = exec.ArtifactOutput(
			"some-plan-id",
			"some-name",
		)

		stepErr = step.Run(ctx, state)
	})

	It("is successful", func() {
		Expect(step.Succeeded()).To(BeTrue())
	})

	Context("when the artifact is present", func() {
		var (
			output *bytes.Buffer
			source *workerfakes.FakeArtifactSource
		)

		BeforeEach(func() {
			output = new(bytes.Buffer)
			source = new(workerfakes.FakeArtifactSource)

			state.Artifacts().RegisterSource("some-name", source)

			go state.ReadPlanOutput("some-plan-id", output)
		})

		It("waits for a user and sends the artifact to them", func() {
			Expect(source.StreamToCallCount()).To(Equal(1))

			dest := source.StreamToArgsForCall(0)
			Expect(dest.StreamIn(".", bytes.NewBufferString("hello"))).To(Succeed())

			Expect(output.String()).To(Equal("hello"))
		})
	})

	Context("when the artifact is not present", func() {
		BeforeEach(func() {
			// do nothing
		})

		It("returns UnknownArtifactSourceError", func() {
			Expect(stepErr).To(Equal(exec.UnknownArtifactSourceError{
				SourceName: "some-name",
			}))
		})
	})
})
