package rpubsub

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"

	"github.com/batchcorp/plumber/dynamic"
	"github.com/batchcorp/plumber/validate"
)

func (r *RedisPubsub) Dynamic(ctx context.Context, dynamicOpts *opts.DynamicOptions) error {
	if err := validateDynamicOptions(dynamicOpts); err != nil {
		return errors.Wrap(err, "invalid dynamic options")
	}

	llog := logrus.WithField("pkg", "rabbitmq/dynamic")

	// Start up dynamic connection
	grpc, err := dynamic.New(dynamicOpts, "Redis PubSub")
	if err != nil {
		return errors.Wrap(err, "could not establish connection to Batch")
	}

	go grpc.Start()

	// Continually loop looking for messages on the channel.
	for {
		select {
		case outbound := <-grpc.OutboundMessageCh:
			for _, ch := range dynamicOpts.RedisPubsub.Args.Channel {
				err := r.client.Publish(context.Background(), ch, outbound.Blob).Err()
				if err != nil {
					llog.Errorf("Unable to replay message: %s", err)
					break
				}

				llog.Debugf("Replayed message to Redis PubSub channel '%s' for replay '%s'", ch, outbound.ReplayId)
			}
		case <-ctx.Done():
			r.log.Warning("context cancelled")
			return nil
		}

	}

	return nil
}

func validateDynamicOptions(dynamicOpts *opts.DynamicOptions) error {
	if dynamicOpts == nil {
		return validate.ErrEmptyDynamicOpts
	}

	if dynamicOpts.RedisPubsub == nil {
		return validate.ErrEmptyBackendGroup
	}

	if dynamicOpts.RedisPubsub.Args == nil {
		return validate.ErrEmptyBackendArgs
	}

	if len(dynamicOpts.RedisPubsub.Args.Channel) == 0 {
		return ErrMissingChannel
	}

	return nil
}